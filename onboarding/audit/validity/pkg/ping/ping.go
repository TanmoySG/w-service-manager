package ping

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	ContentType = "application/json"

	GET  = "get"
	POST = "post"
)

func getCall(uri string, payload *map[string]interface{}) string {
	response, _ := http.Get(uri)
	return response.Status
}

func postCall(uri string, payload *map[string]interface{}) string {
	requestJSON, _ := json.Marshal(payload)
	response, _ := http.Post(uri, ContentType, bytes.NewBuffer(requestJSON))
	return response.Status
}

func Ping(uri string, method string, payload *map[string]interface{}) (string, error) {
	switch strings.ToLower(method) {
	case GET:
		return getCall(uri, payload), nil
	case POST:
		return postCall(uri, payload), nil
	default:
		return "", fmt.Errorf("Error Pinging")
	}
}
