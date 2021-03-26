package bloomfilter

import (
	"hash/crc32"
	"hash/fnv"
)

// BloomFilter interface define
type BloomFilter interface {
	Put([]byte) error
	MightContain([]byte) (bool, error)
}

//hash
func hashFunc(b []byte) uint64 {
	h := fnv.New64a()
	_, _ = h.Write(b)
	return h.Sum64()
}
func hashFunc1(b []byte) uint64 {
	h := fnv.New32()
	_, _ = h.Write(b)
	return uint64(h.Sum32())
}
func hashFunc2(b []byte) uint64 {
	h := crc32.NewIEEE()
	_, _ = h.Write(b)
	return uint64(h.Sum32())
}
