package main

type consistentHashing struct {
	arr  []string
	hash Hash

	ringSize int
}

func NewConsistentHashing(ringSize int) consistentHashing {
	return consistentHashing{
		arr:      make([]string, ringSize),
		hash:     NewHash(),
		ringSize: ringSize,
	}
}

func (ch *consistentHashing) PlotKey(key string) int {
	hashKey := ch.hash.hash(key, ch.ringSize, 1)
	ch.arr[hashKey] = key
	return hashKey
}

func (ch *consistentHashing) PlotServer(name string) int {
	hashKey := ch.hash.hash(name, ch.ringSize, 1)
	ch.arr[hashKey] = name
	return hashKey
}

func (ch *consistentHashing) getHashValue(index int) string {
	return ch.arr[index]
}
