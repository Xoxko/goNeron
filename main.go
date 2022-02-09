package main

import (
	"GONERON/neron"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("LL")
	rand.Seed(time.Now().UnixNano())
	nerv := neron.Neron
	nerv.NewNode(2, 2, 3, 2)
	nerv.RandomNode()
	fmt.Println("nerv.H", len(nerv.NODE[0]))
}
