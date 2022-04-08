<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: mission_generator.proto

namespace Pb\Mission;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>pb.Mission.Twist</code>
 */
class Twist extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string title = 1;</code>
     */
    protected $title = '';
    /**
     * Generated from protobuf field <code>string description = 2;</code>
     */
    protected $description = '';
    /**
     * Generated from protobuf field <code>string rules = 3;</code>
     */
    protected $rules = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $title
     *     @type string $description
     *     @type string $rules
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\MissionGenerator::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string title = 1;</code>
     * @return string
     */
    public function getTitle()
    {
        return $this->title;
    }

    /**
     * Generated from protobuf field <code>string title = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setTitle($var)
    {
        GPBUtil::checkString($var, True);
        $this->title = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string description = 2;</code>
     * @return string
     */
    public function getDescription()
    {
        return $this->description;
    }

    /**
     * Generated from protobuf field <code>string description = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setDescription($var)
    {
        GPBUtil::checkString($var, True);
        $this->description = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string rules = 3;</code>
     * @return string
     */
    public function getRules()
    {
        return $this->rules;
    }

    /**
     * Generated from protobuf field <code>string rules = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setRules($var)
    {
        GPBUtil::checkString($var, True);
        $this->rules = $var;

        return $this;
    }

}

// Adding a class alias for backwards compatibility with the previous class name.
class_alias(Twist::class, \Pb\Mission_Twist::class);

