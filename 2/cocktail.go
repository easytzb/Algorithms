// Cocktail Sort
package main

//import "fmt"

func cocktail(arr []int) {
	left := 0
	right := len(arr) - 1
	//fmt.Println(arr)
	for left < right {
		//from left to right
		for i := left; i < right; i++ {
			if arr[i] > arr[i+1] {
				arr[i] ^= arr[i+1]
				arr[i+1] ^= arr[i]
				arr[i] ^= arr[i+1]
			}
		}

		right--

		//from right to left
		for i := right; i > left; i-- {
			if arr[i] < arr[i-1] {
				arr[i] ^= arr[i-1]
				arr[i-1] ^= arr[i]
				arr[i] ^= arr[i-1]
			}
		}
		left++
	}
	//fmt.Println(arr)
}
