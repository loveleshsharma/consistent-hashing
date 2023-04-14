package main

import "hash/fnv"

type Hash struct {
}

func NewHash() Hash {
	return Hash{}
}

func (h Hash) hash(key string, max, min int) int {
	hs := fnv.New32a()
	hs.Write([]byte(key))
	hash := hs.Sum32()

	rangeSize := max - min + 1
	return int(hash%uint32(rangeSize)) + min
}
