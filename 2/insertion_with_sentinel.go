//Insertion sort
package main

func insertion_with_sentinel(arr []int) {
	len := len(arr)
	min := 0

	for i := 1; i < len; i++ {
		if arr[i] < arr[min] {
			min = i
		}
	}
	if min != 0 {
		tmp := arr[0]
		arr[0] = arr[min]
		arr[min] = tmp
	}

	for i := 1; i < len; i++ {
		for j := i; arr[j] < arr[j-1]; j-- {
			tmp := arr[j]
			arr[j] = arr[j-1]
			arr[j-1] = tmp
		}
	}
}
