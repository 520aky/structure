package main

import "fmt"

//单链表

type HeroNode struct {
	No       int
	Name     string
	Nickname string
	Next     *HeroNode //这个表示指向下一个地址
}

//给链表插入一个新节点（插入顺序排序）
func InsertHeroNode(head *HeroNode, hero *HeroNode) {
	tmp := head
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = hero
}

//有序插入按照No排序
func InsertHeroNodeByNo(head *HeroNode, hero *HeroNode) {
	tmp := head
	for tmp.Next != nil {
		if tmp.Next.No > hero.No {
			break
		}
		tmp = tmp.Next
	}
	hero.Next = tmp.Next
	tmp.Next = hero
}

//删除按照No排序
func DelHeroNodeByNo(head *HeroNode, no int) {
	tmp := head
	for tmp.Next != nil {
		if tmp.Next.No == no {
			tmp.Next = tmp.Next.Next
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
	hero33 := &HeroNode{
		No:       3,
		Name:     "李逵",
		Nickname: "黑旋风",
	}
	_ = hero33
	InsertHeroNodeByNo(head, hero3)
	//ListHeroNode(head)
	InsertHeroNodeByNo(head, hero1)
	//ListHeroNode(head)
	InsertHeroNodeByNo(head, hero2)

	//ListHeroNode(head)

	ListHeroNode(head)

	DelHeroNodeByNo(head, 1)
	ListHeroNode(head)

}
