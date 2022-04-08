<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Pb;

/**
 */
class MissionGeneratorClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \Pb\GenerateMissionRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function Generate(\Pb\GenerateMissionRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/pb.MissionGenerator/Generate',
        $argument,
        ['\Pb\Mission', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \Pb\ListMissionRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\ServerStreamingCall
     */
    public function List(\Pb\ListMissionRequest $argument,
      $metadata = [], $options = []) {
        return $this->_serverStreamRequest('/pb.MissionGenerator/List',
        $argument,
        ['\Pb\Mission', 'decode'],
        $metadata, $options);
    }

}
