package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	// "net"
	"os"
	"time"
)

const (
	BUF_LEN = 1000
)

func main() {
	// StartClient(8081, "./cmd/faketcp/D=2020-05-20_T=18-02-50--air.dat")
	StartClient(8081, "./cmd/faketcp/D=2020-05-20_T=18-03-48--fullbottle05.dat")
	// StartClient(8081, "./cmd/faketcp/room2.dat")
}

func StartClient(port int, filepath string) {
	// Подключаемся к сокету
	// conn, err := net.Dial("tcp", ":"+fmt.Sprint(port))
	// if err != nil {
	// 	fmt.Println("Неудачная попытка подключения")
	// 	os.Exit(1)
	// }

	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	bufSize := make([]byte, 2)

	for {
		reader.Read(bufSize)
		packetSize := binary.BigEndian.Uint16(bufSize)
		buf := make([]byte, packetSize - 0)
		reader.Read(buf)

		fmt.Println(packetSize)

		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}

		// conn.Write(buf)
		time.Sleep(1000 * time.Millisecond)
	}
	fmt.Println("Файл подошел к концу")
}
