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
	_ "github.com/lib/pq"
	_ "github.com/ziutek/mymysql/godrv"
	"log"
	"net/http"
	// "strings"
	"encoding/json"
	// "os"
	"strconv"
	"time"
	"runtime"
)

// /**
var config = readConfig()

// fmt.Println(config)
// var max_pool = strconv.ParseInt(string(config["max_pool_size"]), 10, 0)
// **/

var MAX_POOL_SIZE = 10
var MySQLPool chan *sql.DB

func getPool() *sql.DB {
	config := readConfig()
	parallel_num, _ := strconv.ParseInt(config["parallel_num"], 10, 0) 
	runtime.GOMAXPROCS(int(parallel_num))
	
	dbAdapter := config["db_adapter"]
	dbhost := config["dbhost"]
	dbport := config["dbport"]
	dbuser := config["dbuser"]
	dbpwd := config["dbpwd"]
	dbname := config["dbname"]
	// fmt.Println(config)
	// fmt.Println(config["bind"])
	if MySQLPool == nil {
		MySQLPool = make(chan *sql.DB, MAX_POOL_SIZE)
	}
	
	if len(MySQLPool) == 0 {	//if no DB conn, we make the pool
		max_pool, _ := strconv.ParseInt(config["max_pool_size"], 10, 0)
		// fmt.Println(max_pool)
		go func() {
			for i := 0; i < int(max_pool); i++ {
				fmt.Println("crean DB conn....")
				// mysqlc, err := sql.Open("mymysql", "tcp:127.0.0.1:3306*test/root/")
				
				if dbAdapter == "mysql" {
					//mysql conn
					mysqlc, err := sql.Open("mymysql", "tcp:"+dbhost+":"+dbport+"*"+dbname+"/"+dbuser+"/"+dbpwd)
					if err != nil {
						panic(err)
					}
					putPool(mysqlc)
				} else if dbAdapter == "pgsql" {
					//pgsql conn
					// mysqlc, err := sql.Open("postgres", "user=postgres_user password=password host=172.16.2.29 dbname=my_postgres_db sslmode=disable")
					postgrec, err := sql.Open("postgres", "user="+dbuser+" password="+dbpwd+" host="+dbhost+" port="+dbport+" dbname="+dbname+" sslmode=disable")
					if err != nil {
						panic(err)
					}
					putPool(postgrec)
				}
			}
		}()
	}
	return <-MySQLPool	//parallel return DB conn use pool
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
		/**
		*	the json return back
		* code 0 is for success
		*			 1 is for wrong query
		*			 2 is for wrong get data
		* status success or fail
		* rows query data
		**/
		var back = make(map[string]interface{}) //return json
		back["code"] = 0
		back["status"] = "success"
		back["rows"] = ""

		var errCode = 0
		uukey := r.FormValue("uukey")
		if uukey != config["uukey"] {
			back["code"] = 3
			back["status"] = "fail"
			// log.Fatal("uukey not the same")
			errCode = 1
		}
		
		cache := r.FormValue("cache")	//cache enable
		ccKey := r.FormValue("cache_key")	//cache key
		ccTime := r.FormValue("cctime")	//cache time
		
		if errCode == 0 { //check secure uukey
			query := r.FormValue("query")
			fmt.Println("query is: ", query)
			// os.Exit(3)
			//get cache info
			if query == "cc_info" {
				// ccInfo()
				cc_info := ccInfo()
				fmt.Fprintf(w, cc_info)
				return
			} else if config["cache"] == "1" && query == "cc_get" {	//get cache by key
				// ccGet(ccKey)
			  fmt.Fprintf(w, ccGet(ccKey))	
				return
			}

			//get DB data
			dbx := getPool()	//get the DB conn from pool when we want to use
			rows, errQuery := dbx.Query(query)
			if errQuery != nil {
				back["code"] = 1
				back["status"] = "fail"
				// log.Fatal(err)
				fmt.Println(errQuery)
				errCode = 1
			}

			if errQuery == nil {
				// var email string
				colNames, err := rows.Columns()
				checkErr(err)
				readCols := make([]interface{}, len(colNames))
				writeCols := make([][]byte, len(colNames))
				// writeCols := make([]byte, len(colNames))
				// writeCols := make([]interface{}, len(colNames))
				for i, _ := range writeCols {
					readCols[i] = &writeCols[i]
				}

				result := make([]map[string]interface{}, 0)
				// fmt.Println("len res:", len(result))
				// fmt.Println(result)

				// fmt.Println(colNames)

				for rows.Next() {
					if err := rows.Scan(readCols...); err != nil {
						log.Fatal(err)
						back["code"] = 2
						back["status"] = "fail"
					}

					// fmt.Println(writeCols)
					var tmpStr string
					tmpMap := make(map[string]interface{})

					for i, raw := range writeCols {
						// var tmpStr string
						if raw == nil {
							// result[i] = "\\N"
						} else {
							tmpStr += string(raw)
							tmpMap[colNames[i]] = string(raw)
						}
					}
					result = append(result, tmpMap)
				}

				if err := rows.Err(); err != nil {
					log.Fatal(err)
				}
				// fmt.Println(result)

				// back["rows"] = resStr
				back["rows"] = result
			}
			defer putPool(dbx)	//put the DB conn into pool after finish used
		}
		//just query DB
		jsback, _ := json.Marshal(back)
		
		//if set cache
		if string(config["cache"]) == "1" && cache == "1" && ccKey != "" {
			startTime := strconv.FormatInt(time.Now().Unix(), 10)
			ccJsback := string(jsback) + string(config["cache_split"]) +
									string(startTime) + ":" + ccTime 
			//store cache
			// ccSet(ccKey, string(jsback))
			ccSet(ccKey, string(ccJsback))
		}
		// jsback, _ := json.Marshal(result)
		fmt.Fprintf(w, string(jsback))
		// defer putPool(dbx)
	}
}
