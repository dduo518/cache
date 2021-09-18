package main

import (
	"encoding/binary"
	"io"
	"log"
	"os"
)

type Temp struct {
	file     *os.File
	offset   int64
	fileName string
}

func (temp *Temp) CreateFile() {
	file, err := os.OpenFile(temp.fileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	temp.file = file
	stat, err := os.Stat(temp.fileName)
	if err != nil {
		panic(err)
	}
	temp.offset = stat.Size()
}

func (temp *Temp) Close() {
	_ = temp.file.Close()
}

func (temp *Temp) Write(bs []byte) {
	_, err := temp.file.WriteAt(bs, temp.offset)
	if err != nil {
		panic(err)
	}
	temp.offset += int64(len(bs))
}

func (temp Temp) Read() {
	var offset int64 = 0
	for {
		var entryLenBs = make([]byte, 4)
		_, err := temp.file.ReadAt(entryLenBs, offset)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		offset += 4
		entryLen := binary.BigEndian.Uint32(entryLenBs)
		var entryBs = make([]byte, entryLen-4)
		_, err = temp.file.ReadAt(entryBs, offset)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		offset += int64(entryLen - 4)
		entry := &Entry{}
		entry.Decode(entryBs)
		log.Println(string(entry.key), string(entry.val))
	}
}




