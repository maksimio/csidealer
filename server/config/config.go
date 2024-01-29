package config

import (
	"errors"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// TODO: решить проблему с путями при отладке
const configPath = "/home/m/dev/csidealer/server/config.yml"
const defaultConfigPath = "/home/m/dev/csidealer/server/config/defaultConfig.yml"

func ReadConfig() (*Config, error) {
	config := &Config{}

	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		log.Print("конфигурационный файл не найден")
		err = copyConfig()
		if err != nil {
			log.Fatal("не удалось создать конфигурационный файл по-умолчанию")
		} else {
			log.Print("создан конфигурационный файл с значениями по-умолчанию")
		}
	} else {
		log.Print("найден конфигурационный файл")
	}

	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("не удалось открыть конфигурационный файл %s", configPath)
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		log.Fatalf("не удалось декодировать конфигурационный файл %s", configPath)
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
