//shellBars sort
package main

func shellBars(arr []int) {
	len := len(arr)
	for step := len >> 1; step > 0; step = step >> 1 {
		for i := step; i < len; i++ {
			tmp := arr[i]
			tmp_index := i
			for j := i - step; j >= 0 && arr[j] > tmp; j = j - step {
				arr[j+step] = arr[j]
				tmp_index = j
			}
			if i != tmp_index {
				arr[tmp_index] = tmp
			}
		}
	}
}
