package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sort(f func([]int), times int, length int) {
	t0 := time.Now()
	for i := 0; i < times; i++ {
		arr := rand.Perm(length)
		f(arr)
	}
	t1 := time.Now()
	fmt.Printf("The %v took %v to run.\n", f, t1.Sub(t0))
}

func main() {

	sort(improv_insertion, 1000, 1000)
	sort(insertion, 1000, 1000)
	sort(selection, 1000, 1000)
}
