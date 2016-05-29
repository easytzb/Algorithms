// Bucket Sort
// http://www.cnblogs.com/Braveliu/archive/2013/01/19/2867163.html
// 实现步骤如下：
// <1>求得欲排数据序列中的最大数据。
// <2>通过遍历欲排数据对每个数据乘以10再与最大数据取余，求得每个数据对应桶的索引（或称关键字）。
// 	实际上是将数据分组存入桶中，每个数据*10与最大数据取余即组号；
// 	乘以的数字是10，且是跟最大数据取余，这一步骤决定了最多分成10+1组，证明如下：
// 		e是欲排序序列中的一个数据，eMax是序列中最大值。则e/eMax为[0,1]区间的数，则10*e/eMax为[0,10]区间的数
// 	在大桶号内即余数大的数据一定比余数小的数据要大
//（2）求得每个桶中盛放的数据个数（为了保证随后准确分配数据）
//（3）求得每个桶盛放数据个数的右边界索引（所谓的桶逻辑控制）
//（4）从右向左（确保稳定性）扫描欲排数据序列，依次分配到对应桶中
// 	要保证稳定性，桶内数据排序采用的算法依然要具有稳定性
//（5）对各桶中所盛数据进行收集
//（6）利用插入排序再对各个桶中所盛数据进行排序
//（7）直至排序结束，即为已排序数据序列
package main

//import "fmt"

var bucketCount int = 100

func bucket(arr []int) {
	len := len(arr)

	var max int
	for i := 0; i < len; i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}

	//每个桶存放数据个数
	count := make([]int, bucketCount)
	for i := 0; i < len; i++ {
		count[Map2Index(arr[i], max)]++
	}

	//将个数转化为右边界
	for i := 1; i < bucketCount; i++ {
		count[i] += count[i-1]
	}

	//数据存放桶内
	tmp := make([]int, len)
	for i := len - 1; i > -1; i-- {
		bucketIndex := Map2Index(arr[i], max)
		count[bucketIndex]--
		tmp[(count[bucketIndex])] = arr[i]
	}

	//从桶内移回原数组
	for i := 0; i < len; i++ {
		arr[i] = tmp[i]
	}

	//桶内排序
	//fmt.Println(count)
	for i := 1; i < bucketCount; i++ {
		left, right := count[i-1], count[i]-1
		if left < right {
			insert_sort(arr, left, right)
		}
	}
	//fmt.Println(arr)
}

//决定数据存放的桶号
func Map2Index(val int, max int) int {
	if max == 0 {
		return 0
	}
	return val * (bucketCount - 1) / max
}

//插入排序
func insert_sort(arr []int, start int, end int) {
	min := arr[start]
	minIndex := start
	end++
	for i := start; i < end; i++ {
		if arr[i] < min {
			min = arr[i]
			minIndex = i
		}
	}
	if minIndex != start {
		arr[start] ^= arr[minIndex]
		arr[minIndex] ^= arr[start]
		arr[start] ^= arr[minIndex]
	}

	for i := start + 1; i < end; i++ {
		tmpIndex, tmp := i, arr[i]
		for j := i; arr[j-1] > tmp; j-- {
			arr[j] = arr[j-1]
			tmpIndex = j - 1
		}

		if tmpIndex != i {
			arr[tmpIndex] = tmp
		}
	}
}
