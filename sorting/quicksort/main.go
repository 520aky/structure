package main

import (
	"fmt"
	"time"
)

const arrSize = 6

//快速排序
func QuickSort(left, right int, array *[arrSize]int) {
	l := left
	r := right

	pivot := array[(left+right)/2]
	//for循环的目标是将比pivot小的数放在左边，比pivot大的数放在右边
	for l < r {
		//从pivot左边找到一个大于等于pivot的值
		//重新排序后，会把排序过后的重新检查一遍
		for array[l] < pivot {
			l++
		}
		//从pivot右边找到小于等于pivot的值
		for array[r] > pivot {
			r--
		}
		//表明本次分解任务完成
		if l >= r {
			break
		}
		//交换
		array[l], array[r] = array[r], array[l]
		fmt.Println(array)
		//优化
		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}
		//	fmt.Printf("%v\n", *array)
	}
	//可注释以下代码做测试
	//如果l==r，再移动下，
	//这里注释会报错
	if l == r {
		l++
		r--
	}
	//向左递归
	if left < r {
		QuickSort(left, r, array)
	}
	//向右递归
	if right > l {
		QuickSort(l, right, array)
	}
}

func main() {
	arr := [arrSize]int{90, 61, 60, 10, 20, 50}
	//rand.Seed(time.Now().UnixNano())
	//for i := 0; i < arrSize; i++ {
	//	arr[i] = rand.Intn(arrSize)
	//}

	start := time.Now()
	QuickSort(0, len(arr)-1, &arr)
	elapse := time.Since(start)
	fmt.Println("消耗时间:", elapse, arr)
}
