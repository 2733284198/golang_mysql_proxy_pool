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
  var err error
	if config["https"] != "1" {
		err = http.ListenAndServe(":"+strings.TrimSpace(port), nil)
	} else {
		err = http.ListenAndServeTLS(":"+strings.TrimSpace(port), "server.crt",
			"server.key", nil)
	}
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		fmt.Println(123)
	}
}
