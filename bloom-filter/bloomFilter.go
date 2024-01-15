package bloomfiltergo

import (
	"github.com/spaolacci/murmur3"
)

func generateHash(str string) []uint32 {
	data := []byte(str)
	hashValues := []uint32{}
	for _, v := range SEED {
		hashValues = append(hashValues, murmur3.Sum32WithSeed(data, v))
	}
	return hashValues
}

type bloomFilter struct {
	size   int
	filter []byte
}

func NewBloomFilter(size int) *bloomFilter {
	return &bloomFilter{
		size:   size,
		filter: make([]byte, size),
	}
}

func (b bloomFilter) SetValue(str string) {
	hashValues := generateHash(str)
	for _, valu := range hashValues {
		ind := valu % uint32(b.size)
		b.filter[ind] = 1
	}
}

func (b bloomFilter) GetValuePresent(str string) bool {
	hashValues := generateHash(str)
	resp := true
	for _, valu := range hashValues {
		ind := valu % uint32(b.size)
		resp = resp && (b.filter[ind] == 1)
	}
	return resp
}
