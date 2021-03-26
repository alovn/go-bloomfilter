/*
A simple and high-performance Bloom filter written in golang.

MemoryBloomFilter:

	var bloom BloomFilter = NewMemoryBloomFilter(10000)
	bs := []byte("bloom")
	_ = bloom.Put(bs)
	exists, err := bloom.MightContain(bs)

edisBloomFilter:

	cli := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	key := "redis bloomfilter"
	var bloom BloomFilter = NewRedisBloomFilter(cli, "test", 10000)
	bs := []byte(key)
	_ = bloom.Put(bs)
	exists, err := bloom.MightContain(bs)
*/
package bloomfilter
