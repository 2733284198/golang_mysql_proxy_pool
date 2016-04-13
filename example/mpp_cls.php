<?php 
/**
* Copyright 0x029 Inc. 
* License: MIT License
* Author: JJyy<82049406@qq.com>
* Mpp client PHP class 
**/

class Mpp {
  public $_mppServ; //mpp server addr
  public $_tblPrefix; //table prefix
  protected static $_uukey = "9f89c3e4bee30185648172d9013525cf";
  
  function __construct($mppServ) {
    $this->_mppServ = $mppServ;
    // echo $this->_mppServ;
  }

  /**
  * set table prefix
  *@param: 
  *@return: 
  *
  **/
  function setTblPrefix($prefix) {
    $this->_tblPrefix = $prefix;
  }

  /**
  * 1,fetch data from mpp by query
  * 2,set cache by query
  *
  *@param: string query, array ccArr('key', 'time') 
  *@return: query result
  *
  **/
  function fetch($query, $ccArr=array()) {
    if (empty($ccArr)) {
    $rq = '&query='.$query;
  } else {
    $rq = '&query='.$query.'&cache=1&cache_key='.$ccArr[0].'&cctime='.$ccArr[1];
  }
    $jsonArr = $this->unifyRequest($rq);
    $jsonArr = $this->formatJson($jsonArr);
    return $jsonArr;
  }
  
  /**
  * get cache by key
  *
  *@param:  string key
  *@return: string 
  *
  **/
  function getCache($key) {
    $rq = '&query=cc_get&cache_key='.$key;
    $jsonArr = $this->unifyRequest($rq);
    $jsonArr = $this->formatJson($jsonArr);
    return $jsonArr;
  }
  
  /**
  * set cache by key
  *
  *@param: string key, string val
  *@return: string
  *
  **/
  // function setCache($key, $val) {
  //   $rq = '&query=cc_get&cache_key='.$key;
  //   $jsonArr = $this->unifyRequest($rq);
  //   $jsonArr = $this->formatJson($jsonArr);
  //   return $jsonArr;
  // }
  
  /**
  * get all cache info
  *
  *@param: 
  *@return: 
  **/
  function getAllCc() {
    $rq = '&query=cc_info';
    $jsonArr = $this->unifyRequest($rq);
    return $jsonArr;
  }
  
  
  /**
  * format json string
  *
  *@param:  json string
  *@return: array 
  **/
  function formatJson($jsonStr) {
    $jsonArr = json_decode($jsonStr, true);
    $jsonArr = $jsonArr['rows'];
    return $jsonArr;
  }
  
  /**
  * unify MPP client request code
  *
  *@param:  string $rq 
  *@return: null 
  **/
  function unifyRequest($rq) {
    $cl = curl_init();
    curl_setopt($cl, CURLOPT_URL, $this->_mppServ);
    // echo "query is: ".$rq."\n\n";
    // curl_setopt($cl,CURLOPT_POSTFIELDS,'uukey='.self::$_uukey.'&query='.$query); 
    curl_setopt($cl,CURLOPT_POSTFIELDS,'uukey='.self::$_uukey.$rq); 
    curl_setopt($cl, CURLOPT_RETURNTRANSFER, 1);
    $str = curl_exec($cl);
    curl_close($cl);
    return $str;
  } 
  
}

$mpp = new Mpp("http://127.0.0.1:9090/mpp");
//no set cache, just query DB
// $str = $mpp->fetch("select * from users limit 1");

//set  cache
// $str = $mpp->fetch("select * from users limit 1", array('sql_key', '3600'));

//get cache
$str = $mpp->getCache('sql_key');

//get all cache info
$str = $mpp->getAllCc();

print_r($str);

?>
