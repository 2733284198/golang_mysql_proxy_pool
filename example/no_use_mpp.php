<?php 
/**
* Copyright 0x029 Inc. 
* License: MIT License
* Author: JJyy<82049406@qq.com>
* the sample for no use MPP, direct connect mysql & select data, use mysqli
**/
set_time_limit(0);
$start_time = microtime(true);

$db = new mysqli('127.0.0.1', 'root', '', 'test');
$db->set_charset('utf8');
$query = "select email from users order by id";
if ($res = $db->query($query)) {
  while ($row = $res->fetch_assoc()) {
    echo $row['email'].' ';
  }
  $res->close();
}
$db->close();

$end_time = microtime(true);
$use_time = $end_time - $start_time;
echo "\n-----Use time: ".$use_time."-------\n";

?>