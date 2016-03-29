/**
* Copyright 0x029 Inc. 
* License: MIT License
* Author: JJyy<82049406@qq.com>
* operate cache for MPP, KV format
**/
package main

import (
  "fmt"
	// "encoding/json"
  "time"
  "strings"
  "strconv"
)

// var cc = make(map[string]string) //cache map

func ccSet(key string, val string) {
  cc[key] = val
}

func ccGet(key string) string {
  // return cc[key]
  // /**
  ccStr := cc[key]
  // ccSli := strings.Split(ccStr, "@mpp_split@")
  ccSli := strings.Split(ccStr, config["cache_split"])
  fmt.Println(ccSli)
  ccTimeStr := ccSli[1]
  ccTimeSli := strings.Split(ccTimeStr, ":")
  fmt.Println(ccTimeSli)
  // sTime := int(ccTimeSli[0])
  sTime, _ := strconv.ParseInt(ccTimeSli[0], 10, 64)
  setTime, _ := strconv.ParseInt(ccTimeSli[1], 10, 64)
  timeNow := time.Now().Unix()
  if ((timeNow - sTime) > setTime && setTime != 0) {
    // fmt.Println("cache out of time")
    return "cache out of time"
  } else {
    // fmt.Println(ccSli[0])
    return ccSli[0]
  }
  return "none"
  // **/
}

func ccDel(key string) {
  delete(cc, key)
}


/**
* return all cache info
*
*@author: JJyy
*@param: 
*@return: 
*
**/
func ccInfo() string {
	// info, _ := json.Marshal(cc)
  // return string(info)
  var reStr string
  for k, v := range cc {
    reStr += k+" "+v+"<br><br>"
    // fmt.Println(k, v)
  }
  return reStr
}




