package main

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"runtime"
	"time"
)

func sort(f func([]int), times int, length int) {
	t0 := time.Now()
	for i := 0; i < times; i++ {
		arr := rand.Perm(length)
		f(arr)
	}
	t1 := time.Now()
	fmt.Printf("The %v took %v to run.\n", runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), t1.Sub(t0))
}

func main() {
	funcs := map[string]func([]int){"improv_insertion": improv_insertion, "insertion": insertion, "selection": selection}

	len := len(os.Args)
	if len < 2 {
		fmt.Println("Please special sort algorthms")
		os.Exit(2)
	}

	for i := 1; i < len; i++ {
		if fun, ok := funcs[os.Args[i]]; ok {
			sort(fun, 1, 100)
		}
	}
}
