package main

import (
	"syscall"
	"unsafe"
	"net/http"
	"os"
	"fmt"
	"bytes"
	"encoding/json"
)

var lastInputInfo struct {
    cbSize uint32
    dwTime uint32
}


type Data struct {
	Id string `json:"id"`
	Timestamp string `json:"timestamp"`
}


var (
	user32 = syscall.MustLoadDLL("user32.dll")
	getLastInputInfo = user32.MustFindProc("GetLastInputInfo") 
)


func makeHttpPostReq(url string, id string, timestamp string) {
    client := http.Client{}
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
    lastInputInfo.cbSize = uint32(unsafe.Sizeof(lastInputInfo))
    r1, _, _ := getLastInputInfo.Call(uintptr(unsafe.Pointer(&lastInputInfo)))
    hn, _ := os.Hostname()

    if r1 != 0 {
        makeHttpPostReq("http://localhost:8080/presence", hn, fmt.Sprintf("%v", lastInputInfo.dwTime))
    }
}
