package bloomfilter

import "fmt"

func ExampleNewMemoryBloomFilter() {
	var bloom BloomFilter = NewMemoryBloomFilter(10000)
	bs := []byte("bloom")
	bloom.Put(bs)
	fmt.Println(bloom.MightContain(bs))
}
