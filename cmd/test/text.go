package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	fmt.Println("Testing...")
	buf := []byte{219, 2, 0, 0, 0} // Правильный ответ: 731
	x := binary.LittleEndian.Uint32(buf)
	fmt.Println(x)
}
