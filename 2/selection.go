//Selection sort
package main

func selection(arr []int) {
	len := len(arr)
	for i := 0; i < len; i++ {
		min := i

		for j := i + 1; j < len; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}

		if min != i {
			tmp := arr[i]
			arr[i] = arr[min]
			arr[min] = tmp
		}
	}
}
