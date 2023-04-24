package wsHandlers

import (
	"encoding/json"
	"fmt"
	"tokuteibu/streamer"

	"github.com/gofrs/uuid"
	"github.com/mitchellh/mapstructure"
)

type setViewingRequest struct {
	Id     int `json:"id"`
}

func SetViewing(s *streamer.Streamer, clientId uuid.UUID, args map[string]interface{}) error {
	var req setViewingRequest
	err := mapstructure.Decode(args, &req)
	if err != nil {
		return err
	}
	println(req.Id)

	var res = streamer.Payload{
		Type: "testResponse",
		Body: map[string]interface{}{
			"content": fmt.Sprintf("id: %d", req.Id),
		},
	}

	resJSON, err := json.Marshal(res)
	if err != nil {
		return err
	}

	s.Send(resJSON, func(_ *streamer.Client) bool { return true })
	return nil
}
