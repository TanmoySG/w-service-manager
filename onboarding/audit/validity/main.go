package main

import (
	"validity/pkg/config"

	AuditValidity "validity/internal/audit-validity"

	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetLevel(log.TraceLevel)

	log.Info("Contract Audit-Validation Started")

	c, _ := config.LoadConfigFromFile("./config/secrets/config.secrets.json")

	av := AuditValidity.AuditValidityClient{
		Config:                   *c,
		ServiceDirectory:         c.ServiceConfig.ServiceDirectory,
		ControlList:              c.ServiceConfig.ControlList,
		ContractSourceTopic:      c.ServiceConfig.ContractSourceTopic,
		ValidContractSinkTopic:   c.ServiceConfig.ValidContractSinkTopic,
		InvalidContractSinkTopic: c.ServiceConfig.InvalidContractSinkTopic,
	}
	av.RunAuditValidity()
}
