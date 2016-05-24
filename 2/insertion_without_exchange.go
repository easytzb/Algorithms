// Insertion sort: replace the exchanging with moving
package main

func insertion_without_exchange(arr []int) {
	len := len(arr)
	for i := 1; i < len; i++ {
		k, tmp := i, arr[i]
		for j := i; j > 0 && arr[j-1] > tmp; j-- {
			arr[j] = arr[j-1]
			k = j - 1
		}
		if k != i {
			arr[k] = tmp
		}
	}
}
