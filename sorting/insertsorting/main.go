package main

import (
	"fmt"
	"time"
)

const arrSize = 4

//插入排序  从小到大

func InsertionSort(arr *[arrSize]int) {
	var j int
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		for j = i; j > 0 && temp < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = temp
	}
}

func main() {

	arr := [arrSize]int{}
	//rand.Seed(time.Now().UnixNano())
	//for i := 0; i < arrSize; i++ {
	//	arr[i] = rand.Intn(20)
	//}
	arr = [arrSize]int{6, 5, 3, 1}

	start := time.Now()
	InsertionSort(&arr)
	elapse := time.Since(start)
	fmt.Println("消耗时间:", elapse, arr)

}
