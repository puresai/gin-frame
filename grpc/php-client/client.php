<?php
require __DIR__ . '/vendor/autoload.php';


//用于连接 服务端
$client = new \Hello\GreeterClient('127.0.0.1:50051', [
    'credentials' => Grpc\ChannelCredentials::createInsecure()
]);

//实例化 TestRequest 请求类
$request = new \Hello\HelloRequest();
$request->setName("world");

//调用远程服务
$get = $client->SayHello($request)->wait();

//返回数组
//$status 是数组
list($reply, $status) = $get;

echo $reply->getMessage().PHP_EOL;
// print_r($client->SayHello($request));