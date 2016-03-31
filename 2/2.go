package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var arr []int

	t0 := time.Now()
	for i := 0; i < 10; i++ {
		arr = rand.Perm(1000)
		improv_insertion(arr)
	}
	t1 := time.Now()
	fmt.Printf("The improve_insertion took %v to run.\n", t1.Sub(t0))

	t0 = time.Now()
	for i := 0; i < 10; i++ {
		arr = rand.Perm(1000)
		insertion(arr)
	}
	t1 = time.Now()
	fmt.Printf("The insertion took %v to run.\n", t1.Sub(t0))

	t0 = time.Now()
	for i := 0; i < 10; i++ {
		arr = rand.Perm(1000)
		selection(arr)
	}
	t1 = time.Now()
	fmt.Printf("The selection took %v to run.\n", t1.Sub(t0))
}
