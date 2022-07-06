package main

import (
	"errors"
	"fmt"
	"os"
)

//非环形队列

//使用一个结构体管理队列
type Queue struct {
	maxSize   int
	arrayList [4]int //数组模拟队列
	front     int    //指向数组的最前面 队列头部,不包含队首元素
	rear      int    // 表示指向队列的尾部 包含队尾元素

}

//添加数据到队列
func (this *Queue) AddQueue(val int) error {
	//判断队列是否已满, rear是队列尾部含队列尾部
	if this.rear == this.maxSize-1 {
		return errors.New("queue full")
	}
	this.rear++
	this.arrayList[this.rear] = val
	return nil
}
func (this *Queue) GetQueue() (int, error) {
	//判断队列是否为空
	if this.rear == this.front {
		return -1, errors.New("queue empty")
	}
	this.front++
	return this.arrayList[this.front], nil
}
func (this *Queue) ShowQueue() error {
	if this.rear == this.front {
		fmt.Println("当前队列为空")
		return nil
	}
	fmt.Println("当前队列情况是：")
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, this.arrayList[i])
	}
	fmt.Println()
	return nil
}

func main() {
	queue := &Queue{
		maxSize: 4,
		front:   -1,
		rear:    -1,
	}
	var key string
	var val int
	for {
		fmt.Println("1. 输入add 添加数据到队列")
		fmt.Println("2. 输入get 从队列中获取数据")
		fmt.Println("3. 输入show 显示队列数据")
		fmt.Println("4. 输入exit 退出")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Print("请输入需要添加的数据：")
			fmt.Scanln(&val)
			if err := queue.AddQueue(val); err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("添加成功")
			}
		case "get":
			get, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("从队列中获取数据:%d\n", get)
			}

		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
