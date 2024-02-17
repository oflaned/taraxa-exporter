package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/oflaned/taraxa-exporter/internal/utils"
)

type NodeData struct {
	Synced          bool `json:"synced"`
	SyncingSeconds  int  `json:"syncing_seconds"`
}

func NodeMetrics(url string) (NodeData, error){
    requestData := map[string]interface{}{
        "jsonrpc": "2.0",
        "method":  "get_node_status",
        "params":  []interface{}{},
        "id":      1,
    }

    var metrics NodeData

    requestBody, err := json.Marshal(requestData)
    if err != nil {
        return metrics, err
    }

    resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
    if err != nil {
        return metrics, err
    }
    defer resp.Body.Close()

    responseBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return metrics, err
    }


    err = json.Unmarshal([]byte(responseBody), &metrics)
    if err != nil {
        return metrics, err
    }

    utils.Logger.Info("url: ", url, " Status of sync: ", metrics.Synced)
    return metrics, nil
}