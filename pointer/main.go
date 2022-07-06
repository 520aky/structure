package main

import "fmt"

type Node struct {
	Id   int
	Name string
	Next *Node
}

func ChangeNode(n *Node) {

	n4 := (*Node)(nil)
	_ = n4
	n3 := &Node{
		Id:   3,
		Name: "Mom",
		Next: nil,
	}
	fmt.Printf("ChangeNode 进入函数时  n=%p &n=%p n.id=%p n.name=%p n=%v\n", n, &n, &n.Id, &n.Name, n)
	fmt.Printf("ChangeNode 进入函数时  n3=%p &n3=%p n3.id=%p n3.name=%p n3=%v\n", n3, &n3, &n3.Id, &n3.Name, n3)

	tmp := n
	fmt.Printf("ChangeNode 初始化tmp等于n时  tmp=%p &tmp=%p tmp.id=%p tmp.name=%p tmp=%v\n", tmp, &tmp, &tmp.Id, &tmp.Name, tmp)

	tmp = n3
	fmt.Printf("ChangeNode tmp等于n3时  tmp=%p &tmp=%p tmp.id=%p tmp.name=%p tmp=%v\n", tmp, &tmp, &tmp.Id, &tmp.Name, tmp)

	n.Id = 11111
	fmt.Printf("ChangeNode 改变n的id值  n=%p &n=%p n.id=%p n.name=%p n=%v\n", n, &n, &n.Id, &n.Name, n)

	n = n3
	fmt.Printf("ChangeNode 改变n等于n3  n=%p &n=%p n.id=%p n.name=%p n=%v\n", n, &n, &n.Id, &n.Name, n)
	n.Id = 3333
	fmt.Printf("ChangeNode 改变n等于n3的值  n3=%p &n3=%p n3.id=%p n3.name=%p n3=%v\n", n3, &n3, &n3.Id, &n3.Name, n3)

}

func main2() {
	n1 := &Node{
		Id:   1,
		Name: "Moon",
	}
	n2 := &Node{
		Id:   2,
		Name: "Star",
	}
	n1.Next = n2
	fmt.Printf("初始化 n1=%p &n1=%p n1.id=%p n1.name=%p n1=%v\n", n1, &n1, &n1.Id, &n1.Name, n1)
	fmt.Printf("初始化 n2=%p &n2=%p n2.id=%p n2.name=%p n2=%v\n", n2, &n2, &n2.Id, &n2.Name, n2)
	ChangeNode(n1)

	fmt.Printf("改变后 n1=%p &n1=%p n1.id=%p n1.name=%p n1=%v\n", n1, &n1, &n1.Id, &n1.Name, n1)
	fmt.Printf("改变后 n2=%p &n2=%p n2.id=%p n2.name=%p n2=%v\n", n2, &n2, &n2.Id, &n2.Name, n2)
}

func changeSlice(s1 *[]int, s2 *[]int) {
	fmt.Printf("s1=%p &s2=%p\n", s1, &s1)
	fmt.Printf("s2=%p &s2=%p\n", s2, &s2)
	(*s1)[0] = 11
}

func main() {
	n1 := []int{1, 2, 3, 4}
	n2 := []int{5, 6, 7, 8}
	fmt.Printf("n1=%p &n1=%p\n", n1, &n1)
	fmt.Printf("n2=%p &n2=%p\n", n2, &n2)
	changeSlice(&n1, &n2)
	fmt.Printf("n1=%p &n1=%p n1=%v\n", n1, &n1, n1)
	fmt.Printf("n2=%p &n2=%p n2=%v\n", n2, &n2, n2)
}
