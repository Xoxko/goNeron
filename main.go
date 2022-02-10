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
	for i := range nerv.Input {
		nerv.Input[i] = rand.Float64()
	}
	fmt.Println(nerv.Process())

}
