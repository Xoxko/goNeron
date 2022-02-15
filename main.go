package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/Xoxko/goNeron/neron"
)

func main() {
	fmt.Println("fdfdfd")
	rand.Seed(time.Now().UnixNano())
	nerv := neron.Neron

	a, s := nerv.NewInitNode([]int{1, 4, 3, 5})
	*a[0] = 123
	*s[3] = 231
	nerv.Random(rand.NormFloat64)
	nerv.Calculation(func(x float64) float64 { return math.Tanh(x) })
	fmt.Println(s)
}
