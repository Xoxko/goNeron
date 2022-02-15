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

	a, s := nerv.NewInitNode([]int{1, 4, 3, 5})
	*a[0] = 123
	*s[3] = 231
	nerv.Random(rand.NormFloat64)
	fmt.Println(*a[0])

	x := []int{1, 2, 3}
	y := &x
	fmt.Printf("%T -> %v\n", y, y)
	fmt.Println(s)
	z := []*int{&x[0], &x[1], &x[2]}
	fmt.Printf("%T -> %v\n", z, z)
}
