package main

import "fmt"

//双向链表

type HeroNode struct {
	No       int
	Name     string
	Nickname string
	Pre      *HeroNode //这个指向前一个地址
	Next     *HeroNode //这个表示指向下一个地址
}

//给链表插入一个新节点
func InsertHeroNode(head *HeroNode, hero *HeroNode) {
	tmp := head
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = hero
	hero.Pre = tmp
}

//有序插入按照No排序
func InsertHeroNode2(head *HeroNode, hero *HeroNode) {
	tmp := head
	for tmp.Next != nil {
		if tmp.Next.No > hero.No {
			break
		}
		tmp = tmp.Next
	}
	hero.Next = tmp.Next
	hero.Pre = tmp
	if tmp.Next != nil {
		tmp.Next.Pre = hero
	}
	tmp.Next = hero
}

//删除按照No排序
func DelHeroNodeByNo(head *HeroNode, no int) {
	tmp := head
	for tmp != nil {
		if tmp.No == no {
			if tmp.Next != nil {
				tmp.Next.Pre = tmp.Pre
			}
			tmp.Pre.Next = tmp.Next
			break
		}
		tmp = tmp.Next
	}
}

//显示链表所有节点信息
func ListHeroNode(head *HeroNode) {
	tmp := head
	if tmp.Next == nil {
		fmt.Println("该链表为空")
		return
	}
	fmt.Println("链表节点信息如下:")
	for tmp.Next != nil {
		fmt.Printf("[%d %s %s]===>", tmp.Next.No, tmp.Next.Name, tmp.Next.Nickname)
		tmp = tmp.Next
	}
	fmt.Println()
}

//反向输出
func ListHeroNode2(head *HeroNode) {
	tmp := head
	if tmp.Next == nil {
		fmt.Println("该链表为空")
		return
	}

	for tmp.Next != nil {
		tmp = tmp.Next
	}

	fmt.Println("链表节点信息如下:")
	for tmp.Pre != nil {
		fmt.Printf("[%d %s %s]===>", tmp.No, tmp.Name, tmp.Nickname)
		tmp = tmp.Pre
	}
	fmt.Println()
}

func main() {
	//1. 创建头节点
	head := &HeroNode{}

	//2. 创建新的节点
	hero1 := &HeroNode{
		No:       1,
		Name:     "宋江",
		Nickname: "及时雨",
	}

	hero2 := &HeroNode{
		No:       2,
		Name:     "卢俊义",
		Nickname: "玉麒麟",
	}

	hero3 := &HeroNode{
		No:       3,
		Name:     "林冲",
		Nickname: "豹子头",
	}

	_, _, _ = hero1, hero2, hero3

	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero3)
	InsertHeroNode2(head, hero2)

	ListHeroNode(head)
	fmt.Println()
	fmt.Println("反向输出为：")
	ListHeroNode2(head)
	fmt.Println()

	DelHeroNodeByNo(head, 3)
	fmt.Println("删除后")
	ListHeroNode(head)
	fmt.Println()
	fmt.Println("删除后反向输出为：")
	ListHeroNode2(head)

}
