
golang mysql pool proxy(MPP)
=======
golang开发的一个mysql连接代理池, HTTP返回查询结果。


特性(迭代更新支持)
======
* golang查询mysql， 支持连接池， HTTP返回数据，通过POST传递sql语句 
* 支持多种语言（发送POST请求即可） 
* 读写分离 
* SQL防注入 
* 支持HTTPS 
* 考虑支持分布式HTTP server
* 加入log文件支持 
 
编译启动方式(linux), windows照样
=====
```cd project_path  ```<br />
``` go build  ```<br />
```mv golang_mysql_proxy_pool mpp ```<br />
```./mpp ```

返回数据类似
{"code":0,"rows":"123@123.com","status":"success"} 的 json 格式<br />
<br />
属性说明<br />
    * code 0 is for success<br />
    *	  1 is for wrong query<br />
    *	  2 is for wrong get data <br />
    * status success or fail<br />
    * rows query data <br />


测试
=======
WIN7, go version go1.4.2 windows/386, PHP 5.6, MYSQL5.6, 1万条mysql数据
example目录<br /><br />
没使用MPP的文件，  php no_use_mpp.php  -----Use time: 1.2660720348358-------<br />
使用MPP的文件，  php use_mpp.php   -----Use time: 0.37302088737488-------


[更多请联系<82049406@qq.com>]









