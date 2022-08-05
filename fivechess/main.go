package main

import "fmt"

const (
	ROW = 10
	COL = 10
)

func menu() {
	fmt.Print("\n")
	fmt.Print("*******************************\n")
	fmt.Print("****  欢迎来到五子棋游戏！ ****\n")
	fmt.Print("******    1.进入游戏     ******\n")
	fmt.Print("******    0.退出游戏     ******\n")
	fmt.Print("*******************************\n")

}

//初始化棋盘
func InitBoard(board *[ROW][COL]string, row, col int) {
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			board[i][j] = " "
		}
	}
}

//打印棋盘
func DisplayBoard(board *[ROW][COL]string, row, col int) {
	for i := 0; i < row; i++ {
		fmt.Printf("   %d", i+1) //打印棋盘横坐标提示
	}
	fmt.Print("\n")
	for j := 0; j < col; j++ {
		fmt.Print("---|") //打印首行棋盘
	}
	fmt.Print("\n")

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("%s  |", board[i][j]) //打印竖标
		}
		fmt.Printf("%d ", i+1) //打印棋盘纵坐标提示
		fmt.Print("\n")
		for j := 0; j < col; j++ {
			fmt.Print("---|") //打印横标
		}
		fmt.Print("\n")
	}

}

//玩家走棋
func PlayerMove(board *[ROW][COL]string, row, col int) {
	x, y := 0, 0
	fmt.Print("玩家走:>\n")
	fmt.Printf("请输入坐标(%d:%d)>", row, col)
	for {
		fmt.Scanf("%d %d", &x, &y)
		//判断坐标合法性
		if x >= 1 && x <= row && y >= 1 && y <= col {
			//判断用户输入坐标是否被占用
			if board[x-1][y-1] == " " {
				board[x-1][y-1] = "O"
				break
			} else {
				fmt.Print("该坐标已经被占用\n")
				fmt.Print("请重新输入:>")
				continue
			}
		}
	}
}

//
func isFull(board *[ROW][COL]string, row, col int) bool {
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == " " {
				return false
			}
		}
	}
	return true
}

//判断输赢
func IsWin(board *[ROW][COL]string, row, col int) string {
	i, j := 0, 0
	//检查横向
	for i = 0; i < row; i++ {
		for j = 0; j < col-4; j++ {
			if board[i][j] == board[i][j+1] &&
				board[i][j+1] == board[i][j+2] &&
				board[i][j+2] == board[i][j+3] &&
				board[i][j+3] == board[i][j+4] &&
				board[i][j] != " " {
				return board[i][j]
			}
		}
	}

	//竖线检查
	for i = 0; i < col; i++ {
		for j = 0; j < row-4; j++ {
			if board[j][i] == board[j+1][i] &&
				board[j][i] == board[j+2][i] &&
				board[j][i] == board[j+3][i] &&
				board[j][i] == board[j+4][i] &&
				board[j][i] != " " {
				return board[j][i]
			}
		}
	}

	//斜线检查
	for i = 0; i < row-4; i++ {
		if board[i][i] == board[i+1][i+1] &&
			board[i+1][i+1] == board[i+2][i+2] &&
			board[i+2][i+2] == board[i+3][i+3] &&
			board[i+3][i+3] == board[i+4][i+4] &&
			board[i][i] != " " {
			return board[i][i]
		}
		if board[i][i+4] == board[i+1][i+3] &&
			board[i+1][i+3] == board[i+2][i+2] &&
			board[i+2][i+2] == board[i+3][i+1] &&
			board[i+3][i+1] == board[i+4][i] &&
			board[i][i+4] != " " {
			return board[i][i+4]
		}

	}
	//检查是否平局
	if isFull(board, row, col) {
		return "Q"
	}

	return " "
}

func main() {
	menu()

	var board [ROW][COL]string
	InitBoard(&board, ROW, COL)
	DisplayBoard(&board, ROW, COL)
	var ret string

	for {
		PlayerMove(&board, ROW, COL)
		DisplayBoard(&board, ROW, COL)

		//判断玩家是否赢得游戏
		ret = IsWin(&board, ROW, COL)
		if ret != " " {
			break
		}

	}

	if ret == "O" {
		fmt.Println("玩家赢")
	} else if ret == "Q" {
		fmt.Println("和局")
	} else {
		fmt.Println("电脑赢")
	}
	DisplayBoard(&board, ROW, COL)
}
