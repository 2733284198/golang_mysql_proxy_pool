package main 

import (
  "fmt"
  "time"
  "strconv"
)

func main() {
  now := time.Now().Unix()
  str := "xxx"
  fmt.Println(now)
  fmt.Println(strconv.Itoa(98) + str)
  fmt.Println(strconv.FormatInt(now, 10) + str)
  fmt.Println("hey you")
  fmt.Println("mpp b3 ")
  fmt.Println("fix b3 ")
  fmt.Println("mpp b4")
}
