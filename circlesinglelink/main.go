package main

import "fmt"

type CatNode struct {
	no   int
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, cat *CatNode) {
	//先判断是不是第一只猫
	if head.next == nil {
		head.name = cat.name
		head.no = cat.no
		head.next = head
		//cat.next = head
		return
	}
	tmp := head
	for {
		if tmp.next == head {
			break
		}
		tmp = tmp.next
	}

	tmp.next = cat
	cat.next = head
}

//删除双向链表
func DelCatNode(head *CatNode, no int) *CatNode {
	//判断是否是空链表
	if head.next == nil {
		fmt.Println(`空链表`)
		return head
	}

	tmp := head

	//只有一个节点时
	if head.next == head {
		if head.next.no == no {
			head.next = nil
		} else {
			fmt.Println("没有找到需要删除的节点")
		}
		return head
	}

	flag := false //表示是否找到节点了

	//大于一个节点时
	for {
		if tmp.next.no == no {
			if tmp.next == head {
				head = tmp.next.next
			}
			tmp.next = tmp.next.next
			flag = true
			break
		}

		tmp = tmp.next
		if tmp.next == head {
			//循环完成了,这里会造成当前的tmp没有进行判断，所以在跳出循环的时候，需要重新判断一次
			break
		}
	}

	if !flag {
		if tmp.next.no == no {
			if tmp.next == head {
				head = tmp.next.next
			}
			tmp.next = tmp.next.next
		} else {
			fmt.Println("没有找到需要删除的节点")
		}

	}
	return head
}

func ListCatNode(head *CatNode) {
	if head.next == nil {
		fmt.Println("空链表")
		return
	}
	tmp := head
	for {
		fmt.Printf("[%d %s]-->", tmp.no, tmp.name)
		if tmp.next == head {
			break
		}
		tmp = tmp.next
	}
	fmt.Println()
}

func main() {

	head := &CatNode{}

	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}
	cat2 := &CatNode{
		no:   2,
		name: "tom2",
	}
	cat3 := &CatNode{
		no:   3,
		name: "tom3",
	}

	_, _ = cat2, cat3

	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)
	ListCatNode(head)

	head = DelCatNode(head, 1)
	fmt.Println()
	fmt.Println("删除过后")
	ListCatNode(head)
}
