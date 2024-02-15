package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/oflaned/taraxa-exporter/internal/utils"
)

type nodeData struct {
	Synced          bool `json:"synced"`
	SyncingSeconds  int  `json:"syncing_seconds"`
}


func CheckNodeStatus(url string) (bool, error){
    requestData := map[string]interface{}{
        "jsonrpc": "2.0",
        "method":  "get_node_status",
        "params":  []interface{}{},
        "id":      1,
    }

    requestBody, err := json.Marshal(requestData)
    if err != nil {
        return false, err
	}

    resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
    if err != nil {
        return false, err
    }
    defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return false, err
    }

	var apiData nodeData
	err = json.Unmarshal([]byte(responseBody), &apiData)
    if err != nil {
        return false, err
    }

	utils.Logger.Info("url: ", url, " Status of sync: ", apiData.Synced)
	return apiData.Synced, nil
}