// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    config, err := UnmarshalConfig(bytes)
//    bytes, err = config.Marshal()

package config

import "encoding/json"

func UnmarshalConfig(data []byte) (Config, error) {
	var r Config
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Config) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Config struct {
	Dbconfig      Dbconfig      `json:"dbconfig"`
	KafkaConfig   KafkaConfig   `json:"kafkaConfig"`
	LogConfig     string        `json:"logConfig"`
	Schema        string        `json:"schema"`
	ServiceConfig ServiceConfig `json:"serviceConfig"`
}

type Dbconfig struct {
	Cluster    string `json:"cluster"`
	Collection string `json:"collection"`
	Database   string `json:"database"`
	Token      string `json:"token"`
}

type KafkaConfig struct {
	Cluster  Cluster  `json:"cluster"`
	Consumer Consumer `json:"consumer"`
	Topic    string   `json:"topic"`
}

type Cluster struct {
	Brokers  []string `json:"brokers"`
	ClientID string   `json:"clientId"`
}

type Consumer struct {
	GroupID string `json:"groupId"`
}

type ServiceConfig struct {
	ServiceDirectory         string `json:"serviceDirectory"`
	ControlList              string `json:"controlList"`
	ContractSourceTopic      string `json:"contractSourceTopic"`
	ValidContractSinkTopic   string `json:"validContractSinkTopic"`
	InvalidContractSinkTopic string `json:"invalidContractSinkTopic"`
}
