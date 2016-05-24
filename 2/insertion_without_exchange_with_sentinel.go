// Insertion sort with sentinel: replace the exchanging with moving
package main

func insertion_without_exchange_with_sentinel(arr []int) {
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
		k, tmp := i, arr[i]
		for j := i; arr[j-1] > tmp; j-- {
			arr[j] = arr[j-1]
			k = j - 1
		}
		if k != i {
			arr[k] = tmp
		}
	}
}
