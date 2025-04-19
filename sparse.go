package ctrie

import "math/bits"

func flagPos(hashcode, bitmap uint64, level uint) (uint64, uint64) {
	index := (hashcode >> level) & 0x3f
	flag := uint64(1) << uint64(index)
	mask := uint64(flag - 1)
	pos := uint64(bits.OnesCount64(bitmap & mask))
	return flag, pos
}
