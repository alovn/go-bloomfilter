package main

import (
	"fmt"

	bloomfilter "github.com/alovn/go-bloomfilter"
	"github.com/redis/go-redis/v9"
)

func main() {
	MemoryBloomFilterExample()
	RedisBloomFilterExample()
}

func MemoryBloomFilterExample() {
	bloom := bloomfilter.NewMemoryBloomFilter(10000)
	bs := []byte("bloom")
	_ = bloom.Put(bs)
	exists, err := bloom.MightContain(bs)
	fmt.Println(exists, err)
}

func RedisBloomFilterExample() {
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
	fmt.Println(exists, err)
}
