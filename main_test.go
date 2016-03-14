/**
* Copyright 0x029 Inc. 
* License: MIT License
* Author: JJyy<82049406@qq.com>
* This file is .... 
**/
package main

import (
	"fmt"
	"testing"
	// "strconv"
)

func TestCc(t *testing.T) {
  k := "name"
  v := "jimmy"
  ccSet(k, v)
  getV := ccGet(k) 
	fmt.Println("val is ", getV)
}

