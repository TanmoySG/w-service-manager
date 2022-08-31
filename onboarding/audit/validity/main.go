package main

import (
	// "fmt"
	// "time"

	// "encoding/json"
	// "validity/pkg/kafka"
	"fmt"
	// "validity/internal/config"
	// "validity/internal/wdb"
	"validity/internal/schema"
)

func main() {

	fmt.Print("Placeholder")

	// broker := []string{
	// 	0: "localhost:9092",
	// }
	// clientID := "hsgfsd"
	// deadline := time.Now().Add(10 * time.Second)

	// kc := kafka.Client{
	// 	Brokers:  broker,
	// 	ClientID: clientID,
	// 	ReadDeadline: deadline,
	// }

	// tpc := "intake"
	// msg := make(map[string]interface{})

	// msg["test"] = "pass"
	// msgData, err := json.Marshal(msg)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// kc.Producer(tpc, []byte("1c"), []byte(msgData))

	// kc.Consumer(tpc, func(key, value []byte) {
	// 	fmt.Printf("%s : %s", key, value)
	// })
	// fmt.Print("test")

	// c, _ := config.LoadConfigFromFile("./config/secrets/config.secrets.json")

	// db := c.Dbconfig.Database

	// w := wdb.Client{
	// 	Cluster:  c.Dbconfig.Cluster,
	// 	Token:    c.Dbconfig.Token,
	// 	Database: &db,
	// }

	// collection := "IntakeRequest-Stage"

	// marker := make(map[string]string)
	// marker["Key"] = "requestID"
	// marker["Value"] = "request-2b3f501d-376a-4cf3-8e85-b3c1fe0cc4df"

	// w.GetData(collection, marker, func(resp wdb.ResponseBody, err error) {
	// 	fmt.Println(resp, err)
	// })

	// data := wdb.Data{
	// 	"test" : "1",
	// 	"result" : "pass",
	// }

	// collection = "testSpace"

	// w.AddData(collection, data, func(rd wdb.ResponseBody, err error) {
	// 	fmt.Println(rd, err)
	// } )

	// marker["Key"] = "_id"
	// marker["Value"] = "J8fAtnxVdNPkoxhAne2tzB"

	// data = wdb.Data{
	// 	"result" : "failedddd",
	// }

	// w.UpdateData(collection, marker, data, func(rb wdb.ResponseBody, err error) {
	// 	fmt.Println(rb, err)
	// } )

	// w.DeleteData(collection, marker, func(rb wdb.ResponseBody, err error) {
	// 	fmt.Println(rb, err)
	// } )

	s := schema.SchemaValidator{
		Schema:   schema.LoadFromFile("../../../schema/service-onboarding/validity.audit.schema.json"),
		Document: schema.LoadFromFile("../../../schema/service-onboarding/contract.intake.schema.json"),
	}
	flag, _ := s.Validate()

	fmt.Println(flag)

}
