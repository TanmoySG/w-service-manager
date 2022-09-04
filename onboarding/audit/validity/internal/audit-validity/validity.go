package AuditValidity

import (
	"fmt"
	"time"
	"validity/internal/checks"
	"validity/pkg/config"
	"validity/pkg/kafka"
	"validity/pkg/wdb"
	"validity/spec/contract"
	"validity/spec/validity"
)

var (
	KAFKA_READ_DEADLINE = time.Now().Add(10 * time.Second)
)

const (
	ValidityKind = "validity.audit.service-onboarding"
)

type AuditValidityClient struct {
	Config           config.Config
	ServiceDirectory string
	ControlList      string
	SourceTopic      string // find better name
	SinkTopic        string // find better name
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

	kc.Consumer(avc.SourceTopic, func(uuid, contractBytes []byte) {
		contractParsed, err := contract.UnmarshalContract(contractBytes)
		if err != nil {
			fmt.Println(err)
		}

		c := checks.NewChecksClient(WDBClient, "mockServiceDirectory", avc.ControlList)

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

		kc.Producer(avc.SinkTopic, uuid, validityByte)

		wdbUpdateMarker := make(map[string]string)
		wdbUpdateMarker["Key"] = "requestID"
		wdbUpdateMarker["Value"] = contractParsed.RequestID

		updatedData := wdb.Data{
			"status": "audit-validated",
		}

		WDBClient.UpdateData("IntakeRequest-Stage", wdbUpdateMarker, updatedData, func(rb wdb.ResponseBody, err error) {
			fmt.Println(validityObject.Checks.Valid)
		})
	})
}
