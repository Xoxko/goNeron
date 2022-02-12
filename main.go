package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Xoxko/goNeron/neron"
)

func main() {
	fmt.Println("fdfdfd")
	rand.Seed(time.Now().UnixNano())
	nerv := neron.Neron
	nerv.NewNode(2, 1, 3, 2)
	nerv.RandomNode()

	for I := 0; I < 1000; I++ {
		a := rand.Int31n(1000)
		b := rand.Int31n(1000)
		nerv.Input[0] = float64(a) / 1000
		nerv.Input[1] = float64(b) / 1000

		s := nerv.Process()

		sum := int32(s[0] * 1000000)

		if (a * b) == sum {
			fmt.Println("Успешно обучилась")
			continue
		} else {
			fmt.Printf("a: %v b: %v \n", a, b)
			fmt.Println("sum: ", a*b)
			fmt.Println("Node: ", sum)
			fmt.Println("++++++++++")
		}
	}

}
