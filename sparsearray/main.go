package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//稀疏数组

type ValNode struct {
	Row int
	Col int
	Val int
}

func main() {
	//1先创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑棋
	chessMap[2][3] = 2 //蓝棋

	//2 输出看看原始数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d ", v2)
		}
		fmt.Println()
	}
	//3 转成一个稀疏数组

	var sparseArr []ValNode

	valNode := ValNode{
		Row: 11,
		Col: 11,
		Val: 0,
	}

	sparseArr = append(sparseArr, valNode)

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				//创建一个值节点
				valNode := ValNode{
					Row: i,
					Col: j,
					Val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	fmt.Println("稀疏数组是::::")
	for i, v := range sparseArr {
		fmt.Printf("%d===> %2d %2d %2d\n", i, v.Row, v.Col, v.Val)
	}

	//将稀疏数组存盘 ./chessMap.data
	file, err := os.OpenFile("./chessMap.data", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, v := range sparseArr {
		writer.WriteString(fmt.Sprintf("%d %d %d\r\n", v.Row, v.Col, v.Val))
	}

	writer.Flush()
	//打开 ./chessMap.data 恢复稀疏数组
	chessMapFile, err := os.Open("./chessMap.data")
	if err != nil {
		log.Fatal(err)
	}
	defer chessMapFile.Close()

	var chessMap2 [11][11]int
	reader := bufio.NewReader(chessMapFile)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		data := strings.Split(string(line), " ")
		row, _ := strconv.Atoi(data[0])
		col, _ := strconv.Atoi(data[1])
		val, _ := strconv.Atoi(data[2])
		if row == 11 && col == 11 && val == 0 {
			continue
		}
		fmt.Println("=====>", row, col, val)

		chessMap2[row][col] = val

	}

	//这里演示 使用稀疏数组直接恢复

	//1.1 先创建一个原始数组
	//var chessMap2 [11][11]int
	//
	////1.2  遍历稀疏数组
	//for i, v := range sparseArr {
	//	if i == 0 {
	//		continue
	//	}
	//	chessMap2[v.Row][v.Col] = v.Val
	//}
	fmt.Println("chessMap2:::::")
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d ", v2)
		}
		fmt.Println()
	}

}
