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

	for {
		reader.Read(bufSize)
		bufSize32[1], bufSize32[0] = bufSize[0], bufSize[1]
		bufSize32[2], bufSize32[3] = 0, 0
		buf := make([]byte, binary.BigEndian.Uint16(bufSize))
		io.ReadFull(reader, buf)

		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}

		conn.Write(bufSize32)
		conn.Write(buf)
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Println("Файл подошел к концу")
}