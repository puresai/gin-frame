<?php
require __DIR__ . '/vendor/autoload.php';

class Client extends \Grpc\BaseStub
{
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * rpc SayHello(HelloRequest) returns (HelloReply) {}
     * 方法名尽量和 (gprc 定义 Greeter 服务)的方法一样
     * 用于请求和响应该服务
     */
    public function SayHello(\Hello\HelloRequest $argument){
        // (/hello.Greeter/SayHello) 是请求服务端那个服务和方法，基本和 proto 文件定义一样
        return $this->_simpleRequest('/hello.Greeter/SayHello',
            $argument,
            ['\Hello\HelloReply', 'decode']
            );
    }

}

//用于连接 服务端WzgFNEiB.8q=
$client = new \Client('127.0.0.1:50051', [
    'credentials' => Grpc\ChannelCredentials::createInsecure()
]);

//实例化 TestRequest 请求类
$request = new \Hello\HelloRequest();
$request->setName("fairy");

//调用远程服务
$get = $client->SayHello($request)->wait();

//返回数组
//$reply 是 TestReply 对象
//$status 是数组
list($reply, $status) = $get;

echo $reply->getMessage().PHP_EOL;
// print_r($client->SayHello($request));