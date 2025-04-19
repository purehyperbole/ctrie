package ctrie

import (
	"testing"
)

var global uint64

func BenchmarkFlagPos(b *testing.B) {
	var flag uint64
	bitmaps := make([]uint64, 4096)

	for i := 0; i < b.N; i++ {
		flag, _ = (flagPos(uint64(i), bitmaps[i%len(bitmaps)], 0))
	}

	// avoid NOOP
	global = flag
}
