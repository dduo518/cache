package main

import "encoding/binary"

const (
	HeaderSize = 12
)

// 存储实体
type Entry struct {
	key     []byte
	val     []byte
	keySize uint32
	valSize uint32
}

func (e Entry) Encode() []byte {
	bufLen := int(e.keySize) + int(e.valSize) + HeaderSize
	buf := make([]byte, bufLen)

	binary.BigEndian.PutUint32(buf[0:4], uint32(bufLen))
	binary.BigEndian.PutUint32(buf[4:8], e.keySize)
	binary.BigEndian.PutUint32(buf[8:12], e.valSize)

	copy(buf[HeaderSize:HeaderSize+int(e.keySize)], e.key)
	copy(buf[HeaderSize+int(e.keySize):HeaderSize+int(e.keySize)+int(e.valSize)], e.val)
	return buf
}

func (e *Entry) Decode(bs []byte) {
	e.keySize = binary.BigEndian.Uint32(bs[0:4])
	e.valSize = binary.BigEndian.Uint32(bs[4:8])
	e.key = make([]byte, e.keySize)
	e.val = make([]byte, e.valSize)
	copy(e.key[:], bs[8:(8+e.keySize)])
	copy(e.val[:], bs[(8+e.keySize):(8+e.keySize+e.valSize)])
}
