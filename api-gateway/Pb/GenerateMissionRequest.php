<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: mission_generator.proto

namespace Pb;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>pb.GenerateMissionRequest</code>
 */
class GenerateMissionRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.pb.MissionFormat missionFormat = 1;</code>
     */
    protected $missionFormat = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $missionFormat
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\MissionGenerator::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.pb.MissionFormat missionFormat = 1;</code>
     * @return int
     */
    public function getMissionFormat()
    {
        return $this->missionFormat;
    }

    /**
     * Generated from protobuf field <code>.pb.MissionFormat missionFormat = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setMissionFormat($var)
    {
        GPBUtil::checkEnum($var, \Pb\MissionFormat::class);
        $this->missionFormat = $var;

        return $this;
    }

}
