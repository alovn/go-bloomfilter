# go-bloomfilter

A simple and high-performance Bloom filter written in golang.

Support memory storage and redis.

## Example

### MemoryBloomFilter

```go
var bloom BloomFilter = NewMemoryBloomFilter(10000)
_ = bloom.Put([]byte("bloom"))
exists,err := bloom.MightContain(bs)

```

### RedisBloomFilter

```go
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
```
