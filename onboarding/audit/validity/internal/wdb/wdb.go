package wdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Cluster  string
	Token    string
	Database *string
}

type Payload struct {
	Database   string                  `json:"database"`
	Collection string                  `json:"collection"`
	Marker     *string                 `json:"marker,omitempty"`
	Data       *map[string]interface{} `json:"data,omitempty"`
}

type RequestBody struct {
	Action  string  `json:"action"`
	Payload Payload `json:"payload"`
}

type ResponseBody map[string]interface{}

const (
	ContentType = "application/json"
)

func (w Client) getConnectionURL() string {
	return fmt.Sprintf("https://wdb.tanmoysg.com/connect?cluster=%s&token=%s", w.Cluster, w.Token)
}

func (w Client) GetData(collection string, marker map[string]string, /*callback func()*/ ) {
	connectionURL := w.getConnectionURL()

	markerValue := fmt.Sprintf("%s : %s", marker["Key"], marker["Value"])

	requestBody := RequestBody{
		Action: "view-data",
		Payload: Payload{
			Database:   *w.Database,
			Collection: collection,
			Marker:     &markerValue,
		},
	}

	requestJSON, _ := json.Marshal(requestBody)

	response, _ := http.Post(connectionURL, ContentType, bytes.NewBuffer(requestJSON))

	body, _ := ioutil.ReadAll(response.Body)

	var result ResponseBody
	_ = json.Unmarshal([]byte(body), &result)

	fmt.Println(result)

}
