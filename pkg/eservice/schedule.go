package eservice

import (
	"encoding/json"
	"io"
	"net/http"
)

type response struct {
    Success bool          `json:"success"`
    Message string        `json:"message"`
    Data    []ScheduleDay `json:"data"`
    Code    string        `json:"code"`
}

func Schedule(group_id string) ([]ScheduleDay, error) {
    raw_response, err := http.Get("https://eservice.omsu.ru/schedule/backend/schedule/group/"+group_id)
    if err != nil {
        return nil, err
    }
    body, err := io.ReadAll(raw_response.Body)
    if err != nil {
        return nil, err
    }
    response := response{}
    json.Unmarshal(body, &response)
    return response.Data, nil
}
