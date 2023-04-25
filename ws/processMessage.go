package ws

import (
	"encoding/json"
	"fmt"

	"tokuteibu/streamer"
)


func ProcessMessage(s *streamer.Streamer, data streamer.ReceiveData) error {
	var req streamer.Payload
	err := json.Unmarshal(data.Payload, &req)
	if err != nil {
		return err
	}

	// fmt.Printf("payload: %v\n", data.payload)
	// fmt.Printf("method: %s\n", req.Method)
	// fmt.Printf("args: %v\n", req.Args)

	switch req.Type {
	case "SET_VIEWING":
		SetViewing(s, data.Id, req.Body)
	default:
		fmt.Printf("invalid method")
	}

	return nil
}
