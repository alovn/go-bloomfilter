package bloomfilter

import (
	"hash/crc32"
	"hash/fnv"
)

type BloomFilter interface {
	Put([]byte)
	MightContain([]byte) bool
}

type memeryBloomFilter struct {
	len       uint64
	bits      []uint64
	hashFuncs []func([]byte) uint64
}

func NewByteBloomFilter(size uint64) BloomFilter {
	return &memeryBloomFilter{
		len:       size,
		bits:      make([]uint64, size),
		hashFuncs: []func([]byte) uint64{hashFunc, hashFunc1, hashFunc2},
	}
}

func (filter *memeryBloomFilter) Put(b []byte) {
	for _, f := range filter.hashFuncs {
		val := f(b)
		idx := val % filter.len //array index
		idx2 := idx % 64        //bit index
		filter.bits[idx] = filter.bits[idx] | (1 << idx2)
	}
}

func (filter *memeryBloomFilter) MightContain(b []byte) bool {
	for _, f := range filter.hashFuncs {
		val := f(b)
		idx := val % filter.len
		idx2 := idx % 64
		if filter.bits[idx]&(1<<idx2) == 0 {
			return false
		}
	}
	return true
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
