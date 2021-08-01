package hash_table

import (
	"math/rand"
	"time"
)

type HashTable struct {
	d int
}

func hashCode(x interface{}) int64 {
	return 0
}

func (hb *HashTable) hash(x interface{}) uint64 {
	rand.Seed(time.Now().UnixNano())
	return uint64(rand.Int63()*hashCode(x)) >> (64 - hb.d)
}
