//Selection sort
package main

func selection(arr []int) {
	//arr := [10]int{9, 4, 11, 5, 6, 7, 1, 10, 8, 3}
	len := len(arr)
	//fmt.Println("b", arr)
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
		//fmt.Println(i+1, arr)
	}

	//fmt.Println("e", arr)
}
