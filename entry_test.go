package main

import (
	"fmt"
	"testing"
)

func BenchmarkEntry_Encode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var e = &Entry{
			key: []byte(fmt.Sprintf("key_%v", i)),
			val: []byte(fmt.Sprintf("value_%v", i)),
		}
		e.keySize = uint32(len(e.key))
		e.valSize = uint32(len(e.val))
		_ = e.Encode()
	}
}

