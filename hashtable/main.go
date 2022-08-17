package main

import (
	"fmt"
	"os"
)

type Emp struct {
	Id   int
	Name string
	Next *Emp
}

func (this *Emp) Show() {
	fmt.Printf("id:%d name:%s\n", this.Id, this.Name)
}

type EmpLink struct {
	Head *Emp
}

func (this *EmpLink) FindById(id int) *Emp {
	cur := this.Head
	if cur == nil {
		return nil
	}

	for cur != nil {
		if cur.Id == id {
			return cur
		}
		cur = cur.Next
	}

	return nil
}

func (this *EmpLink) DelById(id int) {
	cur := this.Head
	if cur == nil {
		fmt.Println("删除失败 链表为空")
		return
	}
	//头就是这个数据
	if cur.Id == id {
		this.Head = cur.Next
		fmt.Println("删除成功 数据在表头")
		return
	}

	var isFind bool
	for cur.Next != nil {
		if cur.Next.Id == id {
			cur.Next = cur.Next.Next
			isFind = true
			break
		}
		cur = cur.Next
	}

	if isFind {
		fmt.Println("删除成功")
	} else {
		fmt.Println("删除失败，未查询到雇员")
	}
	return
}

func (this *EmpLink) Show(index int) {
	cur := this.Head
	fmt.Printf("链表【%d】 =>", index)
	if cur == nil {
		fmt.Println("为空")
		return
	}

	for cur != nil {
		fmt.Printf("id:%d name:%s =>", cur.Id, cur.Name)
		cur = cur.Next
	}
	fmt.Println()
}

func (this *EmpLink) Insert(emp *Emp) {
	if this.Head == nil {
		this.Head = emp
		return
	}

	cur := this.Head
	if cur.Id > emp.Id {
		emp.Next = cur
		this.Head = emp
		return
	}
	for cur.Next != nil {
		if cur.Next.Id > emp.Id {
			break
		}
		cur = cur.Next
	}
	emp.Next = cur.Next
	cur.Next = emp
}

type HashTable struct {
	LinkArr [7]EmpLink
}

func (this *HashTable) Insert(emp *Emp) {
	linkIndex := this.hashFunc(emp.Id)
	this.LinkArr[linkIndex].Insert(emp)
}

func (this *HashTable) Show() {
	for i := 0; i < 7; i++ {
		this.LinkArr[i].Show(i)
	}
}

func (this *HashTable) FindById(id int) *Emp {
	linkIndex := this.hashFunc(id)
	return this.LinkArr[linkIndex].FindById(id)
}

func (this *HashTable) DelById(id int) {
	linkIndex := this.hashFunc(id)
	this.LinkArr[linkIndex].DelById(id)
}

func (this *HashTable) hashFunc(id int) int {
	return id % 7
}

func main() {

	var key string
	var id int
	var name string
	var hashtable HashTable

	for {
		fmt.Println("=============雇员系统菜单=============")
		fmt.Println("input 表示添加雇员")
		fmt.Println("show  表示添加雇员")
		fmt.Println("find  表示查找雇员")
		fmt.Println("del   表示删除雇员")
		fmt.Println("exit  表示退出系统")

		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("请输入雇员id")
			fmt.Scanln(&id)
			fmt.Println("请输入雇员姓名")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashtable.Insert(emp)
		case "show":
			hashtable.Show()
		case "find":
			fmt.Println("请输入需要查询的id")
			fmt.Scanln(&id)
			emp := hashtable.FindById(id)
			if emp == nil {
				fmt.Println("未查询到")
			} else {
				emp.Show()
			}
		case "del":
			fmt.Println("请输入需要删除的id")
			fmt.Scanln(&id)
			hashtable.DelById(id)
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入错误")
		}

	}
}
