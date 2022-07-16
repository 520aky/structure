package main

import "fmt"

const arrSize = 10

//快速排序
func QuickSort(left, right int, array *[arrSize]int) {
	l := left
	r := right

	pivot := array[(left+right)/2]
	//for循环的目标是将比pivot小的数放在左边，比pivot大的数放在右边
	for l < r {
		//从pivot左边找到一个大于等于pivot的值,重新排序后，会把排序过后的重新检查一遍
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
		//l++
		//r--
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
	//如果l==r，在移动下
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
	arr := [arrSize]int{9, 0, 2, 6, 7, 4, 3, 8, 1, 5}
	QuickSort(0, len(arr)-1, &arr)
	fmt.Println(arr)
}
