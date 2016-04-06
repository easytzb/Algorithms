package main

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"time"
)

func sort(f func([]int), times int, length int) {
	t0 := time.Now()
	for i := 0; i < times; i++ {
		arr := rand.Perm(length)
		f(arr)
	}
	t1 := time.Now()
	fmt.Printf("The %v sort %v Arr %v times took %v to run.\n", runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), length, times, t1.Sub(t0))
}

// Can be run as follow:
// go run sqrt.go 9
//go install github.com/easytzb/Algorithms/2; ../bin/2 improv_insertion insertion selection 默认三种算法各执行1000，且slice长度为1000
//go install github.com/easytzb/Algorithms/2; ../bin/2 [1000 [1000]] improv_insertion insertion selection
func main() {
	funcs := map[string]func([]int){"improv_insertion": improv_insertion, "insertion": insertion, "selection": selection}

	len := len(os.Args)
	startIndex, times, sliceLen := 1, 1000, 1000

	if len < 2 {
		fmt.Println("see usage comments in file")
		os.Exit(2)
	}

	secondArg, err2 := strconv.Atoi(os.Args[1])
	if err2 == nil {
		startIndex = 2
		times = secondArg
		thirdArg, err3 := strconv.Atoi(os.Args[2])
		if err3 == nil {
			startIndex = 3
			sliceLen = thirdArg
		}
	}

	for i := startIndex; i < len; i++ {
		if fun, ok := funcs[os.Args[i]]; ok {
			sort(fun, times, sliceLen)
		}
	}
}
