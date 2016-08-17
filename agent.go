package main

import (
	"net/http"
	"os"
	"fmt"
	"bytes"
	"time"
	"encoding/json"
        "errors"
        "io"
        "io/ioutil"
)

type Data struct {
	Id string `json:"id"`
	Timestamp string `json:"timestamp"`
}
    
var client = http.Client { Timeout: 30 * time.Second }

func makeHttpPostReq(url string, id string, timestamp string) (err error){
    var data Data
    data.Id = id
    data.Timestamp = timestamp
    d, _ := json.Marshal(&data)
    var jsonStr = []byte(d)
    req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    resp, err1 := client.Do(req)

    if err1 != nil {
	return errors.New(err1.Error())
    }

    io.Copy(ioutil.Discard, resp.Body)
    resp.Body.Close()

    return
}

var lastTimestamp time.Time
var defaultUrl = os.Getenv("KIROS_URL")

func beat() {
    var diff = time.Since(lastTimestamp)
    var serviceUrl string

    if int(diff.Minutes()) < 10  {
	return
    } 

    if(defaultUrl != "") {
	serviceUrl = defaultUrl
    } else {
  	serviceUrl = "http://kirosprime-pfs.rhcloud.com/beats"
    }

    hn, _ := os.Hostname()
    err := makeHttpPostReq(serviceUrl, hn, fmt.Sprintf("%v", time.Now().Format("2006-01-02T15:04:05Z")))

    if err != nil {
	elog.Warning(1, err.Error())
    } else {
       lastTimestamp = time.Now()
    }
}
