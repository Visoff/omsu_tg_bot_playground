package eservice

import (
	"encoding/json"
	"io"
	"net/http"
)

type reponse struct {
    Code    string  `json:"code"`
    Message string  `json:"message"`
    Data    []Group `json:"data"`
}

func Groups() ([]Group, error) {
    resp, err := http.Get("https://eservice.omsu.ru/schedule/backend/dict/groups")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    resp.Body.Close()
    response := reponse{}
    err = json.Unmarshal(body, &response)
    return response.Data, nil
}
