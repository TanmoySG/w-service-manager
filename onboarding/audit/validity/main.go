package main

import (
	// "fmt"
	// "time"

	// "encoding/json"
	// "validity/pkg/kafka"
	"validity/internal/wdb"
	"validity/internal/config"
)

func main() {

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

	c, _ := config.LoadConfigFromFile("./config/secrets/config.secrets.json")

	db := c.Dbconfig.Database
	collection := "IntakeRequest-Stage"

	w := wdb.Client{
		Cluster: c.Dbconfig.Cluster,
		Token: c.Dbconfig.Token,
		Database: &db,
	}

	 marker :=  make(map[string]string)
	marker["Key"] = "requestID" 
	marker["Value"] = "request-2b3f501d-376a-4cf3-8e85-b3c1fe0cc4df"

	w.GetData(collection, marker)

}
