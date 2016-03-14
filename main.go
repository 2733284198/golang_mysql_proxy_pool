/**
* Copyright 0x029 Inc.
* License: MIT License
* Author: JJyy<82049406@qq.com>
* main file of MPP
**/
package main

import (
	"os"
	// "io"
	"bufio"
	"fmt"
	"log"
	"path"
	"strings"
)

const (
	//IP          string = "127.0.0.1"
	IP          string = "0.0.0.0"
	PORT        string = "9090"
	CONFIG_FILE string = "mpp.conf"
	VERSION     string = "0.1"
)


var cc = make(map[string]string) //cache map

func readConfig() (new_config map[string]string) {
	config := make(map[string]string)
	dir, _ := path.Split(os.Args[0])
	os.Chdir(dir)
	// path, _ := os.Getwd()
	// config_file, err := os.Open(path + "/" + CONFIG_FILE) //打开文件
	config_file, err := os.Open(CONFIG_FILE) //打开文件
	defer config_file.Close()
	if err != nil {
		fmt.Println(err)
		fmt.Print("Can not read configuration file. now exit\n")
		os.Exit(0)
	}
	// buff := bufio.NewReader(config_file) //读入缓存
	//读取配置文件
	/**
		for {
			line, err := buff.ReadString('\n') //以'\n'为结束符读入一行
			if err != nil || io.EOF == err {
				break
			}
			rs := []rune(line)
			if string(rs[0:1]) == `#` || len(line) < 3 {
				continue
			}
			str_type := string(rs[0:strings.Index(line, " ")])
			detail := string(rs[strings.Index(line, " ")+1 : len(rs)-1])
	    fmt.Println(str_type)
	    fmt.Println(detail)
			config[str_type] = detail
	    fmt.Println(str_type)
	    fmt.Println(detail)
		}
	  **/

	scanner := bufio.NewScanner(config_file)
	for scanner.Scan() {
		var text = scanner.Text()
		// fmt.Println(text)
		if len(text) != 0 {
			// fmt.Println("xxx")
			// os.Exit(1)
			// }
			if string(text[0:1]) == "#" || len(text) < 3 {
				continue
			}

			str_type := strings.TrimSpace(string(text[0:strings.Index(text, " ")]))
			detail := strings.TrimSpace(string(text[strings.Index(text, " ")+1 : len(text)]))
			// fmt.Println(str_type)
			// fmt.Println(detail)
			config[str_type] = detail
		}
	}

	// fmt.Println(config)
	//再次过滤 (防止没有配置文件)
	return verify(config)
	// return config
}

func readConfigyy() {
	file, err := os.Open(CONFIG_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

//验证配置文件
func verify(config map[string]string) (config_bak map[string]string) {
	if len(config["bind"]) < 3 {
		config["bind"] = IP
	}
	if len(config["port"]) < 1 {
		config["port"] = PORT
	}
	// fmt.Println(config)
	return config
}

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
