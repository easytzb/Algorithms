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

func newtonIterator(start float64) func(n float64) float64 {
	sqr := start
	return func(n float64) float64 {
		sqr = (sqr + n/sqr) / 2.0
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

	n, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	if n < 0.0 {
		fmt.Println("arg should be a Positive integer")
		return
	}

	start := 1.0
	f := newtonIterator(start)

	rightSqrt, lastSqrt, sqrt := math.Sqrt(n), start, 0.0
	for ; math.Abs(lastSqrt-sqrt) > 1.0e-10; sqrt = f(n) {
		fmt.Println(rightSqrt, sqrt, lastSqrt, math.Abs(lastSqrt-sqrt))
		lastSqrt = sqrt
	}
	fmt.Println("last: ", rightSqrt, sqrt)
}
