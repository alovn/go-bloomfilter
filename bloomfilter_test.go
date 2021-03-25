package bloomfilter

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestByteBloomFilter_MightContain(t *testing.T) {
	var bloom BloomFilter = NewByteBloomFilter(10000)
	rand.Seed(time.Hour.Nanoseconds())
	for i := 0; i < 10000; i++ {
		bs := []byte(fmt.Sprintf("%d", rand.Int63()))
		bloom.Put(bs)
		if !bloom.MightContain(bs) {
			t.Errorf("bloomfilter error")
		}
	}
}

func Benchmark_ByteBloomFilter_MightContain(t *testing.B) {
	var bloom BloomFilter = NewByteBloomFilter(10000)
	for i := 0; i < t.N; i++ {
		bs := []byte(fmt.Sprintf("%d", i))
		bloom.Put(bs)
		if !bloom.MightContain(bs) {
			t.Errorf("bloomfilter error")
		}
	}
}
