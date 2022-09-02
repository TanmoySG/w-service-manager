package serviceDirectory

import (
	"encoding/json"
	"fmt"

	"validity/pkg/wdb"
)

type Service struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

type ServicesList []Service

type MockClient struct {
	WDBClient  wdb.Client
	Collection string
}

func (mc MockClient) GetServiceNameList() *[]string {
	var result []string

	mc.WDBClient.GetAllData(mc.Collection, func(rb wdb.ResponseBody, er error) {
		dataBytes, err := json.Marshal(rb["data"])
		if err != nil {
			fmt.Println("error:", err)
		}

		dataList := make(map[string]Service)
		err = json.Unmarshal(dataBytes, &dataList)
		if err != nil {
			fmt.Println("error:", err)
		}

		for _, v := range dataList {
			result = append(result, v.Name)
		}
	})

	return &result
}
