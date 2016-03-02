<?php 
/**
* Copyright 0x029 Inc. 
* License: MIT License
* Author: JJyy<82049406@qq.com>
* 添加测试数据的.
**/
$conn = mysql_connect('127,0,0,1', 'root', '');
mysql_select_db('test');

for ($i=0; $i<3000; $i++) {
  $query = "insert into users set email='123@123.com'";
  mysql_query($query);
}


?>