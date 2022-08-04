package main

import (
	"errors"
	"fmt"
	"strconv"
)

const Size = 20

//运算符栈
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

//判断是不是运算符
func isOper(oper int) bool {
	if oper == 42 || oper == 43 || oper == 45 || oper == 47 {
		return true
	}
	return false
}

//计算
func cal(num1, num2, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num2 * num1
	case 47:
		res = num2 / num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	default:
		fmt.Println("操作符错误")
	}
	return res
}

//编写方法返回运算符优先级
// * / 优先级为1
// + - 优先级为0
func Priority(oper int) int {
	if oper == 42 || oper == 47 {
		return 1
	} else if oper == 43 || oper == 45 {
		return 0
	} else {
		return -1
	}
}

func main() {
	//数字栈
	numStack := &Stack{
		MaxTop: Size,
		Top:    -1,
	}
	//操作符栈
	operStack := &Stack{
		MaxTop: Size,
		Top:    -1,
	}
	exp := "5+2*6-2"

	flag := false
	num1 := 0
	num2 := 0
	oper := 0
	ret := 0
	index := 0
	for {
		ch := exp[index : index+1]
		v := int([]byte(ch)[0])
		//是操作符
		if isOper(v) {
			//如果是一个空栈直接加进去
			if operStack.isEmpty() {
				operStack.Push(v)
				flag = false
			} else {
				//如果是操作符需要从判断优先级
				vPrior := Priority(v)
				topPrior := Priority(operStack.Arr[operStack.Top])
				//优先级
				if topPrior >= vPrior {
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					oper, _ = operStack.Pop()
					ret = cal(num1, num2, oper)
					//将计算结果重新入栈
					numStack.Push(ret)
					//将当前符号入栈
					operStack.Push(v)
					flag = false
				} else {
					operStack.Push(v)
					flag = false
				}
			}
		} else {
			//表示前面已经有数字, 则把上一个数字组合成新数字再压入
			val, _ := strconv.Atoi(ch)
			if flag {
				vv, _ := numStack.Pop()
				vvv := vv*10 + val
				numStack.Push(vvv)
			} else {
				numStack.Push(val)
			}

			flag = true
		}

		index++
		if index == len(exp) {
			break
		}

	}

	for !operStack.isEmpty() {
		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = operStack.Pop()
		ret = cal(num1, num2, oper)
		//将计算结果重新入栈
		numStack.Push(ret)
	}

	//如果算法没有问题 表达式也是正确的 ，则结果 就是numStack中的最后一个数
	result, _ := numStack.Pop()
	fmt.Printf("表达式为%s, 结果为%d\n", exp, result)

}
