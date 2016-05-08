//Bubble Sort
package main

//import "fmt"

func bubble(arr []int) {
	len := len(arr)
	for i := 1; i < len; i++ {
		len2 := len - i

		for j := 0; j < len2; j++ {
			next := j + 1
			if arr[j] > arr[next] {
				arr[j] ^= arr[next]
				arr[next] ^= arr[j]
				arr[j] ^= arr[next]
			}
		}
	}
	//fmt.Println(arr)
}
