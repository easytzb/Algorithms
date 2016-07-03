// Merge Sort
// https://zh.wikipedia.org/wiki/%E5%BD%92%E5%B9%B6%E6%8E%92%E5%BA%8F
package main

//import "fmt"

func merge(arr []int) {
	len := len(arr)
	//fmt.Println(arr)

	arrTmp := make([]int, len)

	for i := 1; i < len; i <<= 1 {
		step := i << 1

		for j := 0; j < len; j += step {
			left, mid, right, end, index := j, j+i, j+i, j+step, j-1

			if right < len {
				for k := 0; k < step && right < end && left < mid && right < len; k++ {
					index = j + k
					if arr[right] < arr[left] {
						arrTmp[index] = arr[right]
						right++
					} else {
						arrTmp[index] = arr[left]
						left++
					}
				}
			}

			for k := left; k < mid && k < len; k++ {
				index++
				arrTmp[index] = arr[k]
			}

			for k := right; k < end && k < len; k++ {
				index++
				arrTmp[index] = arr[k]
			}
		}

		for j := 0; j < len; j++ {
			arr[j] = arrTmp[j]
		}

		//fmt.Println(arr)
	}
}
