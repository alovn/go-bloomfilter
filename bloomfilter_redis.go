package bloomfilter

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisClient interface {
	Pipeline() redis.Pipeliner
}

type redisBloomFilter struct {
	client       RedisClient
	key          string
	bucketCount  uint64
	bucketMaxLen uint64
	hashFuncs    []func([]byte) uint64
}

func NewRedisBloomFilter(redisCli RedisClient, redisKey string, size uint64) BloomFilter {
	if redisKey == "" {
		redisKey = fmt.Sprintf("_bloomfilter_%d", size)
	}
	var bitMaxLen uint64 = 1000000
	if size <= 0 {
		size = bitMaxLen
	}
	bucketCount := size / bitMaxLen
	if size%bitMaxLen > 0 {
		bucketCount += 1
	}
	bucketCount *= 5
	return &redisBloomFilter{
		client:       redisCli,
		key:          redisKey,
		bucketCount:  bucketCount,
		bucketMaxLen: bitMaxLen,
		hashFuncs:    []func([]byte) uint64{hashFunc, hashFunc1, hashFunc2},
	}
}

func (filter *redisBloomFilter) Put(b []byte) error {
	pipe := filter.client.Pipeline()
	ctx := context.Background()
	var val uint64
	var bucketIdx uint64
	var key string
	for i, f := range filter.hashFuncs {
		val = f(b)
		if i == 0 {
			bucketIdx = val % filter.bucketCount //array index
			key = fmt.Sprintf("%s_%d", filter.key, bucketIdx)
		}
		idx := val % filter.bucketMaxLen //bit index
		_ = pipe.SetBit(ctx, key, int64(idx), 1)
	}
	_, err := pipe.Exec(ctx)
	return err
}

func (filter *redisBloomFilter) MightContain(b []byte) (bool, error) {
	pipe := filter.client.Pipeline()
	ctx := context.Background()
	var results []*redis.IntCmd
	var val uint64
	var bucketIdx uint64
	var key string
	for i, f := range filter.hashFuncs {
		val = f(b)
		if i == 0 {
			bucketIdx = val % filter.bucketCount //array index
			key = fmt.Sprintf("%s_%d", filter.key, bucketIdx)
		}
		idx := val % filter.bucketMaxLen //bit index
		cmd := pipe.GetBit(ctx, key, int64(idx))
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
