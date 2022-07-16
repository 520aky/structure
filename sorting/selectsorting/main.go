package main

import "fmt"

//选择排序

//

func main() {
	arr := []int{6, 8, 3, 4, 5, 7, 9, 1, 0, 2}
	var maxIndex int
	var maxValue int
	for i := 0; i < len(arr)-1; i++ {
		maxValue = arr[i]
		maxIndex = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > maxValue {
				maxValue = arr[j]
				maxIndex = j
			}
		}
		if maxIndex != i {
			arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
		}

	}

	fmt.Println(arr)
}
