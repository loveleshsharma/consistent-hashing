package main

import "fmt"

func main() {

	ch := NewConsistentHashing(1000)

	ch.PlotServer("s1")
	ch.PlotServer("s2")
	ch.PlotServer("s3")

	ch.PlotKey("k3")
	ch.PlotKey("k2")
	ch.PlotKey("k1")

	sr1 := ch.GetServer("k1")
	fmt.Println("server for k1: ", sr1.name)

	sr2 := ch.GetServer("k2")
	fmt.Println("server for k2: ", sr2.name)

	sr3 := ch.GetServer("k3")
	fmt.Println("server for k3: ", sr3.name)
}
