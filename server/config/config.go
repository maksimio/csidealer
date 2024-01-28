package config

import (
	"errors"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const configPath = "/home/m/dev/csidealer/server/config.yml" // TODO: решить проблему с путями при отладке
const defaultConfigPath = "/home/m/dev/csidealer/server/config/defaultConfig.yml"

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

func ReadConfig() (*Config, error) {
	config := &Config{}

	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		log.Print("конфигурационный файл не найден")
		err = copyConfig()
		if err != nil {
			return nil, err
		} else {
			log.Print("создан конфигурационный файл с значениями по-умолчанию")
		}
	} else {
		log.Print("найден конфигурационный файл")
	}

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

func copyConfig() error {
	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		log.Println(e.Name())
	}
	data, err := os.ReadFile(defaultConfigPath)
	if err != nil {
		return err
	}
	err = os.WriteFile(configPath, data, 0644)
	return err
}
