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

type Config struct {
	Values FileScheme
	path   string
}

func NewConfig() *Config {
	values := &FileScheme{}

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

	if err := d.Decode(&values); err != nil {
		log.Fatalf("не удалось декодировать конфигурационный файл %s", configPath)
	}

	return &Config{path: configPath, Values: *values}
}

func (c *Config) Update(key string, value interface{}) error {
	data, _ := yaml.Marshal(c.Values)
	var f map[string]interface{}
	yaml.Unmarshal(data, &f) // ошибки быть не может
	// TODO: придумать вариант красивее + чтобы комментарии сохранялись

	f[key] = value

	data, err := yaml.Marshal(f)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &c.Values)
	if err != nil {
		log.Printf("конфиг: ошибка, у поля %s неверный тип", value)
		return err
	}

	err = os.WriteFile(configPath, data, 0)
	if err != nil {
		return err
	}

	log.Printf("конфиг: обновлено поле %s: %s", key, value)

	return nil
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
