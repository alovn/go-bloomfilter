package bloomfilter

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestByteBloomFilter_MightContain(t *testing.T) {
	var bloom BloomFilter = NewMemoryBloomFilter(10000)
	rand.Seed(time.Hour.Nanoseconds())
	for i := 0; i < 10000; i++ {
		key := fmt.Sprintf("%d", rand.Int63())
		bs := []byte(key)
		_ = bloom.Put(bs)
		exists, err := bloom.MightContain(bs)
		if err != nil {
			t.Errorf("bloomfilter error")
		}
		if !exists {
			t.Errorf("key: %s, not exists", key)
		}
	}
}

func Benchmark_ByteBloomFilter_MightContain(t *testing.B) {
	var bloom BloomFilter = NewMemoryBloomFilter(10000)
	for i := 0; i < t.N; i++ {
		key := fmt.Sprintf("%d", i)
		bs := []byte(key)
		_ = bloom.Put(bs)
		exists, err := bloom.MightContain(bs)
		if err != nil {
			t.Errorf("bloomfilter error")
		}
		if !exists {
			t.Errorf("key: %s, not exists", key)
		}
	}
}
