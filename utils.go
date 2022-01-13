package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

// Int2Hex converts an int64 to a byte array
func Int2Hex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}


// ReverseBytes reverses a bytes array
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i,j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}