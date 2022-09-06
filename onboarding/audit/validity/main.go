package main

import (
	"fmt"
	"validity/pkg/config"

	AuditValidity "validity/internal/audit-validity"
)

func main() {

	fmt.Print("Audit Validation Started")

	c, _ := config.LoadConfigFromFile("./config/secrets/config.secrets.json")

	av := AuditValidity.AuditValidityClient{
		Config:                   *c,
		ServiceDirectory:         "",
		ControlList:              "resources/templates/control.list.json",
		SourceTopic:              "intake",
		SinkTopic:                "audit",
		InvalidContractSinkTopic: "invalid",
	}
	av.RunAuditValidity()
}
