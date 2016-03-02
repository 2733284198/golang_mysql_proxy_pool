/**
* Copyright 0x029 Inc. 
* License: MIT License
* Author: JJyy<82049406@qq.com>
* mysql pool setting 
**/
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/ziutek/mymysql/godrv"
	"log"
	"net/http"
	"strings"
	// "os"
)

var MAX_POOL_SIZE = 10
var MySQLPool chan *sql.DB

func getPool() *sql.DB {
	if MySQLPool == nil {
		MySQLPool = make(chan *sql.DB, MAX_POOL_SIZE)
	}

	if len(MySQLPool) == 0 {
		go func() {
			for i := 0; i < MAX_POOL_SIZE; i++ {
				fmt.Println("crean DB conn....")
				mysqlc, err := sql.Open("mymysql", "tcp:127.0.0.1:3306*test/root/")
				if err != nil {
					panic(err)
				}
				putPool(mysqlc)
			}
		}()
	}
	return <-MySQLPool
}

func putPool(conn *sql.DB) {
	// fmt.Println("crean DB conn....")
	if MySQLPool == nil {
		MySQLPool = make(chan *sql.DB, MAX_POOL_SIZE)
	}
	if len(MySQLPool) == MAX_POOL_SIZE {
		conn.Close()
		return
	}
	MySQLPool <- conn
}

func mpp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		query := r.FormValue("query")
		fmt.Println("query is: ", query)
		// os.Exit(3)

		dbx := getPool()
		rows, err := dbx.Query(query)
		if err != nil {
			log.Fatal(err)
		}

		// var email string
    colNames, err := rows.Columns()
    checkErr(err)
    readCols := make([]interface{}, len(colNames))
    writeCols := make([][]byte, len(colNames))
    for i, _ := range writeCols {
      readCols[i] = &writeCols[i]
    }
    
    result := make([]string, len(colNames))
		// fmt.Println(len(result))
    
		for rows.Next() {
			if err := rows.Scan(readCols...); err != nil {
				log.Fatal(err)
			}
      
      for i, raw := range writeCols {
        if raw == nil {
          result[i] = "\\N"
        } else {
          result = append(result, string(raw))
        }
      }
		}
    
		// fmt.Println(len(result))
		// fmt.Println(result)
		result = append(result[1:])
		// fmt.Println(result)
		resStr := strings.Join(result, ", ") 
		fmt.Fprintf(w, resStr)
    
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		defer putPool(dbx)
	}
}











