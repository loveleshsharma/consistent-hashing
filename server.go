package main

type Server struct {
	name string
}

func NewServer(name string) Server {
	return Server{
		name: name,
	}
}
