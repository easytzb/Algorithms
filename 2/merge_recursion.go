// Merge Sort With Recursion
package main

//import "fmt"

func merge_recursion(arr []int) {
	len := len(arr)
	tmp := make([]int, len)
	merge_recursion_func(arr, tmp, 0, len-1)
}

func merge_recursion_func(arr []int, tmp []int, start int, end int) {
	if start >= end {
		return
	}

	mid := (end + start) >> 1
	left, right, i := start, mid+1, start

	//fmt.Println(left, mid)

	merge_recursion_func(arr, tmp, start, mid)
	merge_recursion_func(arr, tmp, right, end)

	for left <= mid && right <= end {
		if arr[left] > arr[right] {
			tmp[i] = arr[right]
			right++
		} else {
			tmp[i] = arr[left]
			left++
		}
		i++
	}

	for left <= mid {
		tmp[i] = arr[left]
		left++
		i++
	}

	for right <= end {
		tmp[i] = arr[right]
		right++
		i++
	}

	for i := start; i <= end; i++ {
		arr[i] = tmp[i]
	}
}
