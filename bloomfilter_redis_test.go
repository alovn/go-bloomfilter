package bloomfilter

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func Test_redisBloomFilter_MightContain(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	defer rdb.Close()
	key := "redis bloomfilter"
	bloom := NewRedisBloomFilter(rdb, "_test", 100)
	_ = bloom.Put([]byte(key))
	exists, err := bloom.MightContain([]byte(key))
	if err != nil {
		t.Error(err)
	}
	if !exists {
		t.Error("key not exists")
	}

	exists, err = bloom.MightContain([]byte(key + "abc"))
	if err != nil {
		t.Error(err)
	}
	if exists {
		t.Error("key not exists")
	}
}

func Test_redisClusterBloomFilter_MightContain(t *testing.T) {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"127.0.0.1:6371",
			"127.0.0.1:6372",
			"127.0.0.1:6373",
		},
		Password: "1234",
	})
	defer rdb.Close()

	bloom := NewRedisBloomFilter(rdb, "_test", 2000000)

	rand := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	key := fmt.Sprintf("redis bloomfilter %d", rand.Intn(10000))

	err := bloom.Put([]byte(key))
	fmt.Println(err)
	exists, err := bloom.MightContain([]byte(key))
	if err != nil {
		t.Error(err)
	}
	if !exists {
		t.Error("key not exists")
	}

	exists, err = bloom.MightContain([]byte(key + "abc"))
	if err != nil {
		t.Error(err)
	}
	if exists {
		t.Error("key not exists")
	}
}

func Benchmark_RedisClusterBloomFilter_MightContain(t *testing.B) {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"127.0.0.1:6371",
			"127.0.0.1:6372",
			"127.0.0.1:6373",
		},
		Password: "1234",
	})
	defer rdb.Close()
	var bloom BloomFilter = NewRedisBloomFilter(rdb, "_test", 2000000)
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

func Benchmark_RedisBloomFilter_MightContain(t *testing.B) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	defer rdb.Close()
	var bloom BloomFilter = NewRedisBloomFilter(rdb, "_test", 2000000)
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
