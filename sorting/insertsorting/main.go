package main

import "fmt"

//插入排序  从小到大

func InsertionSort(arr *[]int) {
	var j int
	for i := 1; i < len(*arr); i++ {
		temp := (*arr)[i]
		for j = i; j > 0 && temp > (*arr)[j-1]; j-- {
			(*arr)[j] = (*arr)[j-1]
		}
		(*arr)[j] = temp
	}
}

func main() {
	arr := []int{6, 8, 3, 4, 5, 7, 9, 1, 0, 2}

	InsertionSort(&arr)

	fmt.Println(arr)
}
