<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Hello;

/**
 * 定义服务
 */
class GreeterClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \Hello\HelloRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function SayHello(\Hello\HelloRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/hello.Greeter/SayHello',
        $argument,
        ['\Hello\HelloReply', 'decode'],
        $metadata, $options);
    }

}
