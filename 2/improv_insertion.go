//Insertion sort: replace the exchanging with moving
package main

import (
	"fmt"
)

func main() {
	arr := [10]int{9, 4, 11, 5, 6, 7, 1, 10, 8, 3}
	len := len(arr)
	fmt.Println("b", arr)
	for i := 1; i < len; i++ {
		tmp := arr[i]
		k := i
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			//tmp := arr[j]
			//arr[j] = arr[j-1]
			//arr[j-1] = tmp
			arr[j] = arr[j-1]
			k = j - 1
		}

		if k != i {
			arr[k] = tmp
		}

		fmt.Println(i, arr)
	}

	fmt.Println("e", arr)
}
