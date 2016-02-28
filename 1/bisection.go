package main

import (
	"fmt"
	"os"
	"strconv"
)

// Can be run as follow:
// go run bisection.go 8, 8 can be replaced with number between 1 and 9
func main() {

	if len(os.Args) != 2 {
		fmt.Println("see usage comments in file")
		return
	}

	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	search, err := strconv.Atoi(os.Args[1])
	if err != nil || search < 1 || search > 9 {
		fmt.Println("see usage comments in file")
		return
	}

	lo := 0
	hi := len(arr) - 1

	for lo <= hi {
		mid := (hi + lo) / 2
		if search > arr[mid] {
			lo = mid - 1
		} else if search < arr[mid] {
			hi = mid + 1
		} else {
			fmt.Println(mid)
			return
		}
	}
}
