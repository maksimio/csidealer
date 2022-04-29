package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

const (
	BUF_LEN = 1000
)

func main() {
	StartClient(8081, "./cmd/faketcp/room2.dat")
}

func StartClient(port int, filepath string) {
	// Подключаемся к сокету
	conn, err := net.Dial("tcp", ":"+fmt.Sprint(port))
	if err != nil {
		fmt.Println("Неудачная попытка подключения")
		os.Exit(1)
	}

	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	buf := make([]byte, BUF_LEN)

	for {
		n, err := reader.Read(buf)

		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}

		conn.Write(buf)
		fmt.Println("Отправлено байт:", n)
		time.Sleep(1000 * time.Millisecond)
	}
	fmt.Println("Файл подошел к концу")
}
