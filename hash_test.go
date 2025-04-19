package ctrie

import (
	"crypto/rand"
	"testing"
)

func BenchmarkHash(b *testing.B) {
	testdata := make([]byte, 1<<23)
	rand.Read(testdata)
	b.ResetTimer()

	var hash uint64

	for i := 0; i < b.N; i++ {
		x := i % (len(testdata) - 32)
		hash = hashfn(testdata[x : x+32])
	}

	_ = hash
}
