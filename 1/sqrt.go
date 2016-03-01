// get sqrt by Newton's method
// https://zh.wikipedia.org/wiki/%E7%89%9B%E9%A1%BF%E6%B3%95
// http://blog.punkid.org/2008/02/28/compute-the-square-root-via-newtons-iteration
// http://jashawn.iteye.com/blog/1659680
// x[k+1] = 1/2(x[k] + N/x[k])

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func newtonIterator(start float64) func(value float64) float64 {
	sqr := start
	return func(value float64) float64 {
		sqr = (sqr + value/sqr) / 2.0
		return sqr
	}
}

// Can be run as follow:
// go run sqrt.go 9
func main() {
	if len(os.Args) != 2 {
		fmt.Println("see usage comments in file")
		return
	}

	num, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	if num < 0 {
		fmt.Println("arg should be a Positive integer")
		return
	}

	sqr := math.Sqrt(num)
	f := newtonIterator(num)
	for i := 0; i < 40; i++ {
		fmt.Println(i, ". ", f(num), ":", sqr)
	}
}
