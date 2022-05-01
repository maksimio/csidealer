package logger

import "fmt"

func RunLoggerInstance(c <-chan bool) {
	for status := range c {
		fmt.Println("Статус:", status)
	}
}