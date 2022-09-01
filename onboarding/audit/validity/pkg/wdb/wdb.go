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

type Data map[string]interface{}

type Payload struct {
	Database   string  `json:"database"`
	Collection string  `json:"collection"`
	Marker     *string `json:"marker,omitempty"`
	Data       *Data   `json:"data,omitempty"`
}

type RequestBody struct {
	Action  string  `json:"action"`
	Payload Payload `json:"payload"`
}

type ResponseBody map[string]interface{}

const (
	ContentType = "application/json"
)

func getError(resp ResponseBody) error {
	switch resp["status_code"].(string) {
	case "0":
		return fmt.Errorf(resp["response"].(string))
	case "1":
		return nil
	default:
		return nil
	}
}

func (w Client) getConnectionURL() string {
	return fmt.Sprintf("https://wdb.tanmoysg.com/connect?cluster=%s&token=%s", w.Cluster, w.Token)
}

func (w Client) GetData(collection string, marker map[string]string, callback func(ResponseBody, error)) {
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

	callback(result, getError(result))
}

func (w Client) AddData(collection string, data Data, callback func(ResponseBody, error)) {
	connectionURL := w.getConnectionURL()

	requestBody := RequestBody{
		Action: "add-data",
		Payload: Payload{
			Database:   *w.Database,
			Collection: collection,
			Data:       &data,
		},
	}

	requestJSON, _ := json.Marshal(requestBody)

	response, _ := http.Post(connectionURL, ContentType, bytes.NewBuffer(requestJSON))

	body, _ := ioutil.ReadAll(response.Body)

	var result ResponseBody
	_ = json.Unmarshal([]byte(body), &result)

	callback(result, getError(result))
}

func (w Client) UpdateData(collection string, marker map[string]string, data Data, callback func(ResponseBody, error)) {
	connectionURL := w.getConnectionURL()

	markerValue := fmt.Sprintf("%s : %s", marker["Key"], marker["Value"])

	requestBody := RequestBody{
		Action: "update-data",
		Payload: Payload{
			Database:   *w.Database,
			Collection: collection,
			Marker:     &markerValue,
			Data:       &data,
		},
	}

	requestJSON, _ := json.Marshal(requestBody)

	response, _ := http.Post(connectionURL, ContentType, bytes.NewBuffer(requestJSON))

	body, _ := ioutil.ReadAll(response.Body)

	var result ResponseBody
	_ = json.Unmarshal([]byte(body), &result)

	callback(result, getError(result))
}

func (w Client) DeleteData(collection string, marker map[string]string, callback func(ResponseBody, error)) {
	connectionURL := w.getConnectionURL()

	markerValue := fmt.Sprintf("%s : %s", marker["Key"], marker["Value"])

	requestBody := RequestBody{
		Action: "delete-data",
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

	callback(result, getError(result))
}

func (w Client) GetAllData(collection string, callback func(ResponseBody, error)) {
	connectionURL := w.getConnectionURL()

	requestBody := RequestBody{
		Action: "get-data",
		Payload: Payload{
			Database:   *w.Database,
			Collection: collection,
		},
	}

	requestJSON, _ := json.Marshal(requestBody)

	response, _ := http.Post(connectionURL, ContentType, bytes.NewBuffer(requestJSON))

	body, _ := ioutil.ReadAll(response.Body)

	var result ResponseBody
	_ = json.Unmarshal([]byte(body), &result)

	callback(result, getError(result))
}
