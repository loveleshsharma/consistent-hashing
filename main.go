package main

import "fmt"

func main() {

	ch := NewConsistentHashing(1000)

	ch.PlotServer(NewServer("s1"))
	ch.PlotServer(NewServer("s2"))
	ch.PlotServer(NewServer("s3"))

	ch.PlotKey(NewKey("k3"))
	ch.PlotKey(NewKey("k2"))
	ch.PlotKey(NewKey("k1"))

	sr1 := ch.GetServer(NewKey("k1"))
	fmt.Println("server for k1: ", sr1.name)

	sr2 := ch.GetServer(NewKey("k2"))
	fmt.Println("server for k2: ", sr2.name)

	sr3 := ch.GetServer(NewKey("k3"))
	fmt.Println("server for k3: ", sr3.name)
}
