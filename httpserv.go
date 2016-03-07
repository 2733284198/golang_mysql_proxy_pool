/**
* Copyright 0x029 Inc. 
* License: MIT License
* Author: JJyy<82049406@qq.com>
* start HTTP server for client 
**/
package main 

import (
  // "database/sql"
  "fmt"
  // "flag"
  "log"
  "net/http"
  "strings"
)

func startHttpServ() {
  http.HandleFunc("/mpp", mpp)
  config := readConfig()
  // var port = ":9090"
  var port = config["port"]
  // fmt.Println(port)
  err := http.ListenAndServe(":"+strings.TrimSpace(port), nil) 
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
    fmt.Println(123)
  }
}



