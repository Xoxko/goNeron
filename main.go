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

	a := rand.Int() % 10
	b := rand.Int() % 10
	nerv.Input[0] = float64(a) / 10
	nerv.Input[1] = float64(b) / 10

	for I := 0; I < 100000; I++ {

		s := nerv.Process()

		sum := int(s[0] * 100)

		if (a * b) == sum {
			fmt.Println("Успешно обучилась")
			fmt.Println("sum: ", a*b)
			fmt.Println("Node: ", sum)

			a = rand.Int() % 10
			b = rand.Int() % 10
			nerv.Input[0] = float64(a) / 100
			nerv.Input[1] = float64(b) / 100

		} else {
			//fmt.Printf("a: %v b: %v \n", a, b)
			//fmt.Println("sum: ", a*b)
			//fmt.Println("Node: ", sum)
			//fmt.Println("++++++++++")
			s[0] = float64(a*b) / 100
			nerv.Training(s, 0.01)
		}
		nerv.Free_summer()
	}

}
