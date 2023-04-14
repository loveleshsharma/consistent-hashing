package main

const defaultRingSize = 1000

type consistentHashing struct {
	arr  []string
	hash Hash
}

func NewConsistentHashing() consistentHashing {
	return consistentHashing{
		arr:  make([]string, defaultRingSize),
		hash: NewHash(),
	}
}

func (ch *consistentHashing) PlotKey(key string) int {
	hashKey := ch.hash.hash(key, 1000, 1)
	ch.arr[hashKey] = key
	return hashKey
}

func (ch *consistentHashing) PlotServer(name string) int {
	hashKey := ch.hash.hash(name, 1000, 1)
	ch.arr[hashKey] = name
	return hashKey
}

func (ch *consistentHashing) getHashValue(index int) string {
	return ch.arr[index]
}
