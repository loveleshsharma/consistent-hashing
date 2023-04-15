package main

type Key struct {
	name string
}

func NewKey(name string) Key {
	return Key{
		name: name,
	}
}
