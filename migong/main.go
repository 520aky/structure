package main

import "fmt"

//编写一个函数完成老鼠找出路
//myMap 地图
//i j 当前对地图哪个点做测试

func SetWay(myMap *[8][7]int, i, j int) bool {
	//
	if myMap[6][5] == 2 {
		return true
	} else {
		//判断当前点是否探测过
		if myMap[i][j] == 0 {
			//如果为0 表示此路是通的，
			myMap[i][j] = 2 //假定此路是通的，继续递归探测,
			//探测路径为下右上左
			if SetWay(myMap, i+1, j) { //下
				return true
			} else if SetWay(myMap, i, j+1) { //右
				return true
			} else if SetWay(myMap, i-1, j) { //上
				return true
			} else if SetWay(myMap, i, j-1) { //左
				return true
			} else {
				//走到这里，说明此路不通
				myMap[i][j] = 3
				return false
			}
			//
		} else {
			//如果探测过或者不通 ，直接false表示此路不通
			return false
		}
	}
}

func main() {
	//二维数组模拟迷宫
	//1. 如果元素的值为1， 就表示墙
	//2. 元素的值为0 代表这条路没有探测过（没有走过）
	//3. 元素的值为2 表示是可以走的路
	//4. 元素值为3 表示曾经走过，但是走不通
	var mmap [8][7]int

	//初始化墙,先把地图最上最下设置位1
	for i := 0; i < 7; i++ {
		mmap[0][i] = 1
		mmap[7][i] = 1
	}

	//初始化左边和右边的墙
	for i := 0; i < 8; i++ {
		mmap[i][0] = 1
		mmap[i][6] = 1
	}

	mmap[3][1] = 1
	mmap[3][2] = 1

	//输出地图
	for _, v := range mmap {
		fmt.Println(v)
	}

	SetWay(&mmap, 1, 1)

	//输出地图
	for _, v := range mmap {
		fmt.Println(v)
	}
}
