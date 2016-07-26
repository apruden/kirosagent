package main

import (
	"net/http"
	"os"
	"fmt"
	"bytes"
	"time"
	"encoding/json"
)


type Data struct {
	Id string `json:"id"`
	Timestamp string `json:"timestamp"`
}


func makeHttpPostReq(url string, id string, timestamp string) {
    client := http.Client{ Timeout: 1 * time.Second}
    var data Data
    data.Id = id
    data.Timestamp = timestamp
    d, _ := json.Marshal(&data)
    var jsonStr = []byte(d)
    req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    client.Do(req)
}


func beep() {
    hn, _ := os.Hostname()
    makeHttpPostReq("http://localhost:8080/presence", hn, fmt.Sprintf("%v", time.Now().Unix()))
}
