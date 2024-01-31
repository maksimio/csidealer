package main

import (
	"csidealer/config"
	"csidealer/internal/app"
	"fmt"
	"log"
)

func main() {
	log.Println("сервер начинает работу")
	config, err := config.ReadConfig()
	if err != nil {
		fmt.Println("выход из программы - возникла ошибка при чтении конфига:", err)
		return
	}
	app.Run(config)
}
