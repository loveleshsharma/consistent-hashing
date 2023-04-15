package main

type consistentHashing struct {
	arr  []interface{}
	hash Hash

	ringSize int
}

func NewConsistentHashing(ringSize int) consistentHashing {
	return consistentHashing{
		arr:      make([]interface{}, ringSize),
		hash:     NewHash(),
		ringSize: ringSize,
	}
}

func (ch *consistentHashing) PlotKey(key string) int {
	hashKey := ch.hash.hash(key, ch.ringSize, 1)
	ch.arr[hashKey] = NewKey(key)
	return hashKey
}

func (ch *consistentHashing) PlotServer(name string) int {
	hashKey := ch.hash.hash(name, ch.ringSize, 1)
	ch.arr[hashKey] = NewServer(name)
	return hashKey
}

func (ch *consistentHashing) getHashValue(index int) interface{} {
	return ch.arr[index]
}
