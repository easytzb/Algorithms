//Insertion sort
package main

func insertion(arr []int) {
	len := len(arr)
	for i := 1; i < len; i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			tmp := arr[j]
			arr[j] = arr[j-1]
			arr[j-1] = tmp
		}
	}
}
