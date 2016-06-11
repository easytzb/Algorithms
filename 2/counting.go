// Counting Sort
// https://zh.wikipedia.org/wiki/%E8%AE%A1%E6%95%B0%E6%8E%92%E5%BA%8F
// 实现步骤如下：
// <1>找出待排序的数组A中最大和最小的元素
//	其目的在于得出A中不重复元素的最大个数，也就是极值差加1，也是求各额外数组的大小
// <2>创建额外数组C，C的大小为最大元素-最小元素+1
// <3>统计数组中值为i的元素的个数，并存入C中的第(i-最小值)个元素
// <3>求得A中每种元素在最终排序中的右边界，即遍历C，C[i] = C[i] + C[i-1]
// <4>创建新的数组b，大小与A相同
// <5>倒序遍历A，根据C将A中的数据填充到B；B即所求
//	这里倒序遍历，是为了保证算法的稳定性。<3>中得到的是右边界，所以必须倒序，若是左边界，则可以顺序遍历
package main

//import "fmt"

func counting(arr []int) {
	len, min, max := len(arr), 0, 0

	//求最大最小值
	for i := 0; i < len; i++ {
		if min > arr[i] {
			min = arr[i]
		}

		if max < arr[i] {
			max = arr[i]
		}
	}

	countSize := max - min + 1
	count := make([]int, countSize)

	//个数统计
	for i := 0; i < len; i++ {
		count[arr[i]-min]++
	}

	//右边界
	for i := 1; i < countSize; i++ {
		count[i] += count[i-1]
	}

	tmp := make([]int, len)
	for i := len - 1; i > -1; i-- {
		count[arr[i]-min]--
		tmp[count[arr[i]-min]] = arr[i]
	}
	arr = tmp
	//fmt.Println(arr)
}
