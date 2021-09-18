package main

import (
	"fmt"
	"testing"
)

func BenchmarkTemp_Write(b *testing.B) {
	var temp = &Temp{
		fileName: "./temp/test.txt",
	}
	temp.CreateFile()
	defer temp.Close()
	for i := 0; i < 10000; i++ {
		temp.Write([]byte("benchmark"))
	}
}

func BenchmarkTemp_Write2(b *testing.B) {
	var temp = &Temp{
		fileName: "./temp/test.txt",
	}
	temp.CreateFile()
	defer temp.Close()
	for i := 0; i < 10000; i++ {
		var e = &Entry{
			key: []byte(fmt.Sprintf("key_%v", i)),
			val: []byte(fmt.Sprintf("value_%v", i)),
		}
		e.keySize = uint32(len(e.key))
		e.valSize = uint32(len(e.val))
		buf := e.Encode()
		temp.Write(buf)
	}
}