
golang mysql pool proxy(MPP)
=======
mysql procy pool by golang, select query and use HTTP return data

<a href="README_cn.md">英文说明最详细，简单的可以看中文说明</a>

Future(coming soon...)
======
* golang connect mysql， connect pool， HTTP get data，use POST for sql query 
* support postgre database 
* support many program language (just use POST query) 
* write, read separate 
* deal SQL inject
* support HTTPS 
* distributed HTTP server
* support log file
 
build & run (linux), windows the same
=====
```cd project_path  ```<br />
```go build  ```<br />
```mv golang_mysql_proxy_pool mpp ```<br />
```./mpp ``` <br /><br />

#### how run as daemon & log file
if you want to run as daemon, you can do like this(for linux):

```nohub  ./mpp > mpp.log 2>&1 & ```

that will output the log info to the file mpp.log, more info just man nohub 


#### how to use cache ? 
the client for PHP in the folder ./example/mpp_cls.php, the sample code 
for every function of MPP client, please check it

Test
=======
WIN7, go version go1.4.2 windows/386, PHP 5.6, MYSQL5.6, 10000 mysql rows
example folder<br /><br />
dont use MPP php file，  php no_use_mpp.php  -----Use time: 1.2660720348358-------<br />
use MPP php file，  php use_mpp.php   -----Use time: 0.37302088737488-------


[more contact<82049406@qq.com>]









