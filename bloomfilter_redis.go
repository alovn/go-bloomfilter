package bloomfilter

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type redisBloomFilter struct {
	client       *redis.Client
	key          string
	bucketCount  uint64
	bucketMaxLen uint64
	hashFuncs    []func([]byte) uint64
}

func NewRedisBloomFilter(redisCli *redis.Client, redisKey string, size uint64) BloomFilter {
	if redisKey == "" {
		redisKey = fmt.Sprintf("_bloomfilter_%d", size)
	}
	var bucketMaxLen uint64 = 1000000
	bucketCount := size / bucketMaxLen
	if size%bucketMaxLen > 0 {
		bucketCount += 1
	}
	return &redisBloomFilter{
		client:       redisCli,
		key:          redisKey,
		bucketCount:  bucketCount,
		bucketMaxLen: bucketMaxLen,
		hashFuncs:    []func([]byte) uint64{hashFunc, hashFunc1, hashFunc2},
	}
}

func (filter *redisBloomFilter) Put(b []byte) error {
	pipe := filter.client.Pipeline()
	ctx := context.TODO()
	for _, f := range filter.hashFuncs {
		val := f(b)
		bucketIdx := val % filter.bucketCount //array index
		idx2 := val % filter.bucketMaxLen     //bit index
		pipe.SetBit(ctx, fmt.Sprintf("%s_%d", filter.key, bucketIdx), int64(idx2), 1)
	}
	_, err := pipe.Exec(ctx)
	return err
}
func (filter *redisBloomFilter) MightContain(b []byte) (bool, error) {
	pipe := filter.client.Pipeline()
	ctx := context.TODO()
	var results []*redis.IntCmd
	for _, f := range filter.hashFuncs {
		val := f(b)
		bucketIdx := val % filter.bucketCount //array index
		idx2 := val % filter.bucketMaxLen     //bit index
		cmd := pipe.GetBit(ctx, fmt.Sprintf("%s_%d", filter.key, bucketIdx), int64(idx2))
		results = append(results, cmd)
	}
	if _, err := pipe.Exec(ctx); err != nil {
		return false, err
	}
	for _, res := range results {
		bit, err := res.Result()
		if err != nil {
			return false, err
		}
		if bit == 0 {
			return false, nil
		}
	}
	return true, nil
}
