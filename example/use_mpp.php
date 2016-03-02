<?php 
/**
* Copyright 0x029 Inc. 
* License: MIT License
* Author: JJyy<82049406@qq.com>
* the sample for use MPP to connect mysql & select data 
* use CURL for post http data
**/
$start_time = microtime(true);
$cl = curl_init();
curl_setopt($cl, CURLOPT_URL, "http://127.0.0.1:9090/mpp");
curl_setopt($cl,CURLOPT_POSTFIELDS,'query=select email from users order by id'); 
curl_setopt($cl, CURLOPT_RETURNTRANSFER, 1);
$str = curl_exec($cl);
echo $str;
curl_close($cl);
$end_time = microtime(true);

$use_time = $end_time - $start_time;
echo "\n-----Use time: ".$use_time."-------\n";
?>