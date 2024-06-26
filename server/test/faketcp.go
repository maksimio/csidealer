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
	RunTcpWriter(8081, "./test/example.dat")
}

func RunTcpWriter(port int, filepath string) {
	conn, err := net.Dial("tcp", ":"+fmt.Sprint(port))
	if err != nil {
		log.Print(err)
		return
	}
	defer conn.Close()
	log.Println("успешное подключение к серверу на порту", port)

	bufSize := make([]byte, 2)
	bufSize32 := make([]byte, 4)

	for {
		f, err := os.Open(filepath)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		reader := bufio.NewReader(f)
		i := 0

		for {
			reader.Read(bufSize)
			bufSize32[1], bufSize32[0] = bufSize[0], bufSize[1]
			bufSize32[2], bufSize32[3] = 0, 0
			buf := make([]byte, binary.BigEndian.Uint16(bufSize))
			_, err := io.ReadFull(reader, buf)

			if err != nil {
				if err != io.EOF {
					log.Fatal(err)
				}
				break
			}

			conn.Write(bufSize32)
			conn.Write(buf)
			log.Print(i)
			time.Sleep(100 * time.Millisecond)
			i += 1
		}
		log.Print("файл закончился. Повторное чтение")
		time.Sleep(3 * time.Second)
	}
}
