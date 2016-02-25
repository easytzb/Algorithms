// Compute the greatest common divisor of
// two nonnegative integers p and q as follows:
// If q is 0, the answer is p. If not, divide p by q
// and take the remainder r. The answer is the
// greatest common divisor of q and r.

package main

import (
	"fmt"
	"os"
	"strconv"
)

// Can be run as follow:
// go run gcd.go 4 8
func main() {
	if len(os.Args) != 3 {
		fmt.Println("see usage coments in file")
		return
	}

	p, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	q, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(gcd(p, q))
}

func gcd(p int, q int) int {
	if q == 0 {
		return p
	} else {
		return gcd(q, p%q)
	}
}
