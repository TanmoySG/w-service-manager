package AuditValidity

import (
	"fmt"
	"strings"
	"time"
	"validity/internal/checks"
	"validity/pkg/config"
	"validity/pkg/kafka"
	"validity/pkg/wdb"
	"validity/spec/contract"
	"validity/spec/validity"

	log "github.com/sirupsen/logrus"
)

var (
	KAFKA_READ_DEADLINE = time.Now().Add(10 * time.Second)
)

const (
	ValidityKind = "validity.audit.service-onboarding"
)

type AuditValidityClient struct {
	Config                   config.Config
	ServiceDirectory         string
	ControlList              string
	ContractSourceTopic      string // find better name
	ValidContractSinkTopic   string // find better name
	InvalidContractSinkTopic string
}

func getInvalidityLogStatement(contractValidity validity.Checks) string {
	var invalidity []string

	if !contractValidity.DataAccess.Valid {
		invalidity = append(invalidity, contractValidity.DataAccess.Error)
	}

	if !contractValidity.Repository.Valid {
		invalidity = append(invalidity, contractValidity.Repository.Error)
	}

	if !contractValidity.ServiceName.Valid {
		invalidity = append(invalidity, contractValidity.ServiceName.Error)
	}

	return strings.Join(invalidity[:], " , ")
}

func (avc AuditValidityClient) RunAuditValidity() {
	kc := kafka.Client{
		Brokers:      avc.Config.KafkaConfig.Cluster.Brokers,
		ClientID:     avc.Config.KafkaConfig.Consumer.GroupID,
		ReadDeadline: KAFKA_READ_DEADLINE,
	}

	WDBClient := wdb.Client{
		Cluster:  avc.Config.Dbconfig.Cluster,
		Token:    avc.Config.Dbconfig.Token,
		Database: &avc.Config.Dbconfig.Database,
	}

	kc.Consumer(avc.ContractSourceTopic, func(uuid, contractBytes []byte) {
		contractParsed, err := contract.UnmarshalContract(contractBytes)
		if err != nil {
			fmt.Println(err)
		}

		log.Debugf("Proccessing Contract with Contract ID %s", uuid)

		c := checks.NewChecksClient(WDBClient, avc.ServiceDirectory, avc.ControlList)

		contractValidity := c.GetContractValidity(contractParsed)

		validityObject := validity.Validity{
			Checks:      contractValidity,
			Contract:    contractParsed,
			ContractID:  string(uuid),
			Kind:        ValidityKind,
			RequestID:   contractParsed.RequestID,
			ServiceName: contractParsed.Service.Name,
		}

		validityByte, err := validityObject.Marshal()
		if err != nil {
			fmt.Println(err)
		}

		if *contractValidity.Valid {
			kc.Producer(avc.ValidContractSinkTopic, uuid, validityByte)
			log.Infof("Contract with ID %s is %s", uuid, *checks.Valid.Message)
		} else {
			kc.Producer(avc.InvalidContractSinkTopic, uuid, validityByte)
			log.Errorf("Contract with ID %s is Invalid : %v", uuid, getInvalidityLogStatement(contractValidity))
		}

		wdbUpdateMarker := make(map[string]string)
		wdbUpdateMarker["Key"] = "requestID"
		wdbUpdateMarker["Value"] = contractParsed.RequestID

		updatedData := wdb.Data{
			"status": "audit-validation",
		}

		WDBClient.UpdateData(avc.Config.Dbconfig.Collection, wdbUpdateMarker, updatedData, func(rb wdb.ResponseBody, err error) {})
	})
}
