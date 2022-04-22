/*
A simple and high-performance Bloom filter written in golang.

import bloomfilter "github.com/alovn/go-bloomfilter"

MemoryBloomFilter:

	bloom := bloomfilter.NewMemoryBloomFilter(10000)
	bs := []byte("bloom")
	_ = bloom.Put(bs)
	exists, err := bloom.MightContain(bs)

RedisBloomFilter:

	cli := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	key := "redis bloomfilter"
	bloom := bloomfilter.NewRedisBloomFilter(cli, "test", 10000)
	bs := []byte(key)
	_ = bloom.Put(bs)
	exists, err := bloom.MightContain(bs)
*/
package bloomfilter
