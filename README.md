# go-bloomfilter

A simple and high-performance Bloom filter written in golang.

Support memory storage and redis.

## Examples

### MemoryBloomFilter

```go
import bloomfilter "github.com/alovn/go-bloomfilter"

bloom := bloomfilter.NewMemoryBloomFilter(10000)
_ = bloom.Put([]byte("bloom"))
exists,err := bloom.MightContain(bs)

```

### RedisBloomFilter

```go
import (
    bloomfilter "github.com/alovn/go-bloomfilter"
    "github.com/redis/go-redis/v9"
)

rdb := redis.NewClient(&redis.Options{
    Addr:     "127.0.0.1:6379",
    Password: "",
    DB:       0,
})
defer rdb.Close()

key := "redis bloomfilter"

bloom := bloomfilter.NewRedisBloomFilter(rdb, "test", 1000000)

bs := []byte(key)
_ = bloom.Put(bs)
exists, err := bloom.MightContain(bs)
```

redis cluster:

```go
rdb := redis.NewClusterClient(&redis.ClusterOptions{
    Addrs: []string{
        "127.0.0.1:6371",
        "127.0.0.1:6372",
        "127.0.0.1:6373",
    },
    Password: "1234",
})
bloom := bloomfilter.NewRedisBloomFilter(rdb, "test", 1000000)
```
