/*
A simple and high-performance Bloom filter written in golang.

Useage:

	var bloom BloomFilter = NewMemoryBloomFilter(10000)
	bs := []byte("bloom")
	bloom.Put(bs)
	fmt.Println(bloom.MightContain(bs))
*/
package bloomfilter
