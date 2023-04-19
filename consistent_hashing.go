package main

import "fmt"

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

func (ch *consistentHashing) GetServer(key Key) Server {
	hashKey := ch.getHashKey(key.name)

	var isRingVisitCompleted = false

	var i = hashKey
	var condition = ch.ringSize

	//key is at the last index
	if hashKey == ch.ringSize-1 {
		i = 0
		condition = hashKey

		for ; i < condition; i++ {
			if server, ok := ch.arr[i].(Server); ok {
				fmt.Printf("found server: %s\n", server.name)
				return server
			}

		}
	} else {
		for ; i < condition; i++ {
			if server, ok := ch.arr[i].(Server); ok {
				fmt.Printf("found server: %s\n", server.name)
				return server
			}

			if i == condition-1 && isRingVisitCompleted == false {
				isRingVisitCompleted = true
				i = 0
				condition = hashKey
				continue
			}
		}
	}

	fmt.Printf("no server was found in the ring!\n")
	return Server{}
}

func (ch *consistentHashing) RemoveServer(server Server) {
	hashKey := ch.getHashKey(server.name)
	ch.arr[hashKey] = nil
}

func (ch *consistentHashing) PlotKey(key Key) int {
	hashKey := ch.getHashKey(key.name)
	ch.arr[hashKey] = key
	return hashKey
}

func (ch *consistentHashing) PlotServer(server Server) int {
	hashKey := ch.getHashKey(server.name)
	ch.arr[hashKey] = server
	return hashKey
}

func (ch *consistentHashing) getHashValue(index int) interface{} {
	return ch.arr[index]
}

func (ch *consistentHashing) getHashKey(key string) int {
	return ch.hash.hash(key, ch.ringSize)
}
