package main

import (
	"fmt"
	. "shujujiegou/tree/btree"
)

func main() {
	root := NewBSTree(5)
	AddBSTree(NewBSTree(8), root)
	AddBSTree(NewBSTree(2), root)
	AddBSTree(NewBSTree(3), root)
	AddBSTree(NewBSTree(7), root)
	AddBSTree(NewBSTree(3), root)
	AddBSTree(NewBSTree(6), root)
	AddBSTree(NewBSTree(20), root)
	level := root.GetLevel()
	PrintBSTree(BSTrees{root}, level, 1)

	result1 := make([]int, 0)
	root.Preorder(&result1)

	result2 := make([]int, 0)
	root.Inorder(&result2)

	result3 := make([]int, 0)
	root.Postorder(&result3)
	fmt.Println(result1, result2, result3)

	fmt.Printf("最小值为:%d\n", MinNode(root))
	fmt.Printf("最大值为:%d\n", MaxNode(root))

	fmt.Println(SearchNode(7, root))
	fmt.Println("查找父节点:", SearchParentNode(SearchNode(20, root), root))
}
