package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	RunTcpWriter(8081, "./test/data/1.dat")
}

func RunTcpWriter(port int, filepath string) {
	conn, err := net.Dial("tcp", ":"+fmt.Sprint(port))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	fmt.Println("Успешное подключение к серверу на порту", port)

	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bufSize := make([]byte, 2)
	bufSize32 := make([]byte, 4)
	reader := bufio.NewReader(f)
	i := 0

	for {
		reader.Read(bufSize)
		// fmt.Println(bufSize)
		bufSize32[1], bufSize32[0] = bufSize[0], bufSize[1]
		bufSize32[2], bufSize32[3] = 0, 0
		buf := make([]byte, binary.BigEndian.Uint16(bufSize))
		_, err := io.ReadFull(reader, buf)

		if err != nil {
			fmt.Println("Ошибка")
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}

		conn.Write(bufSize32)
		conn.Write(buf)
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(i)
		i += 1
	}
	fmt.Println("Файл подошел к концу")
}
