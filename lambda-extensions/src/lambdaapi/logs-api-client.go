package lambdaapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

const (
	// Base URL for extension
	logsURL = "2020-08-15/logs"
	// Subscription Body Constants. Subscribe to platform logs and receive them on ${local_ip}:4243 via HTTP protocol.
	timeoutMs = 1000
	maxBytes  = 262144
	maxItems  = 1000
)

var (
	logEvents = []EventType{"platform", "function", "extension"}
)

// SubscribeToLogsAPI is - Subscribe to Logs API to receive the Lambda Logs.
func (client *Client) SubscribeToLogsAPI(ctx context.Context) ([]byte, error) {
	URL := client.baseURL + logsURL

	reqBody, err := json.Marshal(map[string]interface{}{
		"destination": map[string]interface{}{"protocol": "HTTP", "URI": fmt.Sprintf("http://sandbox:%v", ReceiverPort)},
		"types":       logEvents,
		"buffering":   map[string]interface{}{"timeoutMs": timeoutMs, "maxBytes": maxBytes, "maxItems": maxItems},
	})
	if err != nil {
		return nil, err
	}
	headers := map[string]string{
		extensionIdentiferHeader: client.extensionID,
	}
	var response []byte
	if ctx != nil {
		response, err = client.MakeRequestWithContext(ctx, headers, bytes.NewBuffer(reqBody), "PUT", URL)
	} else {
		response, err = client.MakeRequest(headers, bytes.NewBuffer(reqBody), "PUT", URL)
	}
	if err != nil {
		return nil, err
	}

	return response, nil
}
