package main

import (
	"fmt"
	"myhash/cmd"
)

func main() {
	//
	hashRing := cmd.New(3, nil)

	//
	hashRing.AddNodes("node1")
	hashRing.AddNodes("node2")
	hashRing.AddNodes("node3")

	//
	fmt.Println(hashRing.GetNode("node1"))
	fmt.Println(hashRing.GetNode("node2"))
	fmt.Println(hashRing.GetNode("node3"))

}
