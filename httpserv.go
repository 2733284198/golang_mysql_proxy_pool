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
  // _ "github.com/ziutek/mymysql/godrv"
  // _ "github.com/go-sql-driver/mysql"
  "log"
  "net/http"
)


func startHttpServ() {
  http.HandleFunc("/mpp", mpp)
  err := http.ListenAndServe(":9090", nil) 
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
    fmt.Println(123)
  }
}



