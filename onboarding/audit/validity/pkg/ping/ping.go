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

func getCall(uri string, payload *map[string]interface{}) int {
	response, _ := http.Get(uri)
	return response.StatusCode
}

func postCall(uri string, payload *map[string]interface{}) int {
	requestJSON, _ := json.Marshal(payload)
	response, _ := http.Post(uri, ContentType, bytes.NewBuffer(requestJSON))
	return response.StatusCode
}

func Ping(uri string, method string, payload *map[string]interface{}) (int, error) {
	switch strings.ToLower(method) {
	case GET:
		return getCall(uri, payload), nil
	case POST:
		return postCall(uri, payload), nil
	default:
		return 0, fmt.Errorf("error pinging")
	}
}
