/**
* Copyright 0x029 Inc. 
* License: MIT License
* Author: JJyy<82049406@qq.com>
* main file of MPP 
**/
package main 

import (
  // "database/sql"
  // "fmt"
  // _ "github.com/ziutek/mymysql/godrv"
  // _ "github.com/go-sql-driver/mysql"
  // "log"
  // "net/http"
)

func checkErr(err error) {
  if err != nil {
    panic(err)
    // log.Fatal(err)
    // fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
    // os.Exit(1)
  }
}

func main() {
  startHttpServ()
}









