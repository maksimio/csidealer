package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Tcp struct {
		Port int `yaml:"port"`
	} `yaml:"tcp"`
	Rx struct {
		Ip       string `yaml:"ip"`
		TargetIp string `yaml:"targetIp"`
	} `yaml:"rx"`
	Tx struct {
		Ip                string `yaml:"ip"`
		IfName            string `yaml:"ifName"`
		DstMacAddr        string `yaml:"dstMacAddr"`
		NumOfPacketToSend int    `yaml:"numOfPacketToSend"`
		PktIntervalUs     int    `yaml:"pktIntervalUs"`
		PktLen            int    `yaml:"pktLen"`
	} `yaml:"tx"`
}

func NewConfig() (*Config, error) {
	configPath := "./config.yml"
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
