# go-bloomfilter

A simple and high-performance Bloom filter written in golang.

## Useage

```go
var bloom BloomFilter = NewMemoryBloomFilter(10000)
bloom.Put([]byte("bloom"))
fmt.Println(bloom.MightContain(bs))
```
