package main

import (
	"errors"
	"fmt"
	"unicode"
)

const Size = 5

//使用数组模拟一个栈的使用
type Stack struct {
	MaxTop int       //表示栈最大可以存放的个数
	Top    int       //表示栈顶，因为栈底是固定的，因此可以不声明
	Arr    [Size]int //数组模拟栈
}

func (this *Stack) Push(val int) error {
	if this.isFull() {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	this.Top++
	this.Arr[this.Top] = val
	return nil
}

func (this *Stack) Pop() (int, error) {
	if this.isEmpty() {
		fmt.Println("empty empty")
		return -1, errors.New("stack empty")
	}
	val := this.Arr[this.Top]
	this.Top--
	return val, nil
}

func (this *Stack) isFull() bool {
	if this.Top == this.MaxTop-1 {
		return true
	}
	return false
}
func (this *Stack) isEmpty() bool {
	if this.Top == -1 {
		return true
	}
	return false
}

//遍历栈 需要从栈顶遍历
func (this *Stack) Show() {
	if this.isEmpty() {
		fmt.Println("stack empty")
		return
	}
	fmt.Println("栈的情况如下:")
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]->%d ", i, this.Arr[i])
	}
	fmt.Print("\n")

}

func main() {
	stack := &Stack{
		MaxTop: Size, //表示栈内最多存放数量
		Top:    -1,   //当栈顶为-1时，表示栈为空
	}
	stack.Push(5)
	stack.Push(6)
	stack.Push(7)
	stack.Push(8)
	stack.Push(9)

	stack.Show()

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	stack.Show()

	//判断一个字符串是个数字
	unicode.IsNumber('3')
}
