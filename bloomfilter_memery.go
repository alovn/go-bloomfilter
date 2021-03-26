package bloomfilter

type memeryBloomFilter struct {
	len       uint64
	bits      []uint64
	hashFuncs []func([]byte) uint64
}

// NewMemoryBloomFilter implementation with memory
func NewMemoryBloomFilter(size uint64) BloomFilter {
	return &memeryBloomFilter{
		len:       size,
		bits:      make([]uint64, size),
		hashFuncs: []func([]byte) uint64{hashFunc, hashFunc1, hashFunc2},
	}
}

func (filter *memeryBloomFilter) Put(b []byte) error {
	for _, f := range filter.hashFuncs {
		val := f(b)
		idx := val % filter.len //array index
		idx2 := val % 64        //bit index
		filter.bits[idx] = filter.bits[idx] | (1 << idx2)
	}
	return nil
}

func (filter *memeryBloomFilter) MightContain(b []byte) (bool, error) {
	for _, f := range filter.hashFuncs {
		val := f(b)
		idx := val % filter.len
		idx2 := val % 64
		if filter.bits[idx]&(1<<idx2) == 0 {
			return false, nil
		}
	}
	return true, nil
}
