/**
* Copyright 0x029 Inc. 
* License: MIT License
* Author: JJyy<82049406@qq.com>
* operate cache for MPP, KV format
**/
package main

import (
  // "fmt"
	// "encoding/json"
)

// var cc = make(map[string]string) //cache map

func ccSet(key string, val string) {
  cc[key] = val
}

func ccGet(key string) string {
  return cc[key]
}

func ccDel(key string) {
  delete(cc, key)
}

// func ccInfo() map[string]interface{} {
func ccInfo() string {
// func ccInfo() {
	// info, _ := json.Marshal(cc)
  // return string(info)
  var reStr string
  for k, v := range cc {
    reStr += k+" "+v+"<br><br>"
    // fmt.Println(k, v)
  }
  return reStr
}

