package main

import (
	"errors"
	"fmt"
	"os"
)

//环形队列,用数组模拟环形队列， 实际队列容量为为maxSize-1，

//使用一个结构体管理队列
type CircleQueue struct {
	maxSize   int
	arrayList [10]int //数组模拟队列
	head      int     //指向数组的最前面 队列头部,不包含队首元素
	tail      int     // 表示指向队列的尾部 包含队尾元素

}

func (this *CircleQueue) IsFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}
func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head
}

//取出环形队列元素个数
func (this *CircleQueue) Size() int {
	//这是一个关键的算法
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

//添加数据到队列
func (this *CircleQueue) Push(val int) error {
	//判断队列是否已满, rear是队列尾部含队列尾部
	if this.IsFull() {
		return errors.New("queue full")
	}
	this.arrayList[this.tail] = val
	this.tail = (this.tail + 1) % this.maxSize
	return nil
}
func (this *CircleQueue) Pop() (int, error) {
	//判断队列是否为空
	if this.IsEmpty() {
		return -1, errors.New("queue empty")
	}
	val := this.arrayList[this.head]
	this.head = (this.head + 1) % this.maxSize
	return val, nil
}

//显示队列
func (this *CircleQueue) ShowQueue() error {
	if this.IsEmpty() {
		fmt.Println("当前队列为空")
		return nil
	}
	fmt.Println("当前队列情况是：")
	//for i := this.head; i <= this.tail; i++ {
	//	fmt.Printf("array[%d]=%d\t", i, this.arrayList[i])
	//}
	head := this.head
	tail := this.tail
	for head != tail {
		fmt.Printf("array[%d]=%d\t", head, this.arrayList[head])
		head = (head + 1) % this.maxSize
	}
	fmt.Println()

	return nil
}

func (this *CircleQueue) ShowQueue2() error {
	size := this.Size()
	if size == 0 {
		fmt.Println("当前队列为空")
		return nil
	}
	fmt.Println("当前队列情况是：")

	head := this.head

	for i := 0; i < size; i++ {
		fmt.Printf("array[%d]=%d\t", head, this.arrayList[head])
		head = (head + 1) % this.maxSize
	}

	fmt.Println()

	return nil
}

func main() {
	queue := &CircleQueue{
		maxSize: 4,
		head:    0,
		tail:    0,
	}
	var key string
	var val int
	fmt.Println("1. 输入push 添加数据到队列")
	fmt.Println("2. 输入pop 从队列中获取数据")
	fmt.Println("3. 输入show 显示队列数据")
	fmt.Println("4. 输入size 显示队列元素个数")
	fmt.Println("5. 输入exit 退出")
	for {
		fmt.Scanln(&key)
		switch key {
		case "push":
			fmt.Print("请输入需要添加的数据：")
			fmt.Scanln(&val)
			if err := queue.Push(val); err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("添加成功")
			}
		case "pop":
			get, err := queue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("从队列中获取数据:%d\n", get)
			}

		case "show":
			queue.ShowQueue2()
		case "size":
			fmt.Printf("队列共有%d个元素\n", queue.Size())
			queue.Size()
		case "exit":
			os.Exit(0)
		}
	}
}
