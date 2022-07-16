package btree

import (
	"fmt"
	"math"
)

//二叉查找树

type BSTree struct {
	Value int
	Left  *BSTree
	Right *BSTree
}

func NewBSTree(value int) *BSTree {
	return &BSTree{Value: value}
}

//添加二叉查找树
func AddBSTree(bst *BSTree, node *BSTree) *BSTree {
	if node == nil {
		return bst
	}
	if bst.Value > node.Value {
		node.Right = AddBSTree(bst, node.Right)
	} else if bst.Value < node.Value {
		node.Left = AddBSTree(bst, node.Left)
	}
	//二叉查找树没有相等的情况, 相等的情况下就返回原来的值
	return node
}

//获取树的高度
func (this *BSTree) GetLevel() int {
	if this == nil {
		return 0
	}

	return max(this.Left.GetLevel(), this.Right.GetLevel()) + 1
}

//先序遍历
func (this *BSTree) Preorder(result *[]int) {
	if this == nil {
		return
	}
	*result = append(*result, this.Value)
	this.Left.Preorder(result)
	this.Right.Preorder(result)
}

//中序遍历
func (this *BSTree) Inorder(result *[]int) {
	if this == nil {
		return
	}

	this.Left.Inorder(result)
	*result = append(*result, this.Value)
	this.Right.Inorder(result)
}

//后序遍历
func (this *BSTree) Postorder(result *[]int) {
	if this == nil {
		return
	}

	this.Left.Postorder(result)
	this.Right.Postorder(result)
	*result = append(*result, this.Value)
}

//判断是否叶节点(就是没有子节点)
func (this *BSTree) isLeaf() bool {
	if this.Left == nil && this.Right == nil {
		return true
	}
	return false
}

//获取最小值
func MinNode(node *BSTree) *BSTree {
	if node.Left == nil {
		return node
	} else {
		return MinNode(node.Left)
	}
}

func MaxNode(node *BSTree) int {
	if node.Right == nil {
		return node.Value
	} else {
		return MaxNode(node.Right)
	}
}

//查找节点
func SearchNode(value int, node *BSTree) *BSTree {
	if node == nil {
		return nil
	}
	if value > node.Value {
		return SearchNode(value, node.Right)
	} else if value < node.Value {
		return SearchNode(value, node.Left)
	} else {
		return node
	}
}

//获取当前节点的子节点
func getSingleChild(node *BSTree) *BSTree {
	if node.Left != nil && node.Right == nil {
		return node.Left
	} else if node.Left == nil && node.Right != nil {
		return node.Right
	} else {
		return nil
	}
}

//查找节点并反问父节点
func SearchNodeWithParent(value int, node *BSTree, parent ...interface{}) (*BSTree, *BSTree, string) {
	if node == nil {
		return nil, nil, ""
	}
	if value > node.Value {
		return SearchNodeWithParent(value, node.Right, node, "right")
	} else if value < node.Value {
		return SearchNodeWithParent(value, node.Left, node, "left")
	} else {
		if len(parent) == 0 {
			return node, nil, ""
		}
		return node, parent[0].(*BSTree), parent[1].(string)
	}
}

//查找父节点
func SearchParentNode(node *BSTree, root *BSTree) *BSTree {
	if node == nil || root == nil || node == root {
		return nil
	}
	//如果相等表示找到
	if node == root.Left || node == root.Right {
		return root
	}
	//递归往左边查找
	left := SearchParentNode(node, root.Left)
	if left != nil {
		return left
	}
	//如果没有往右边查找
	return SearchParentNode(node, root.Right)
}

func DelNode(value int, treeRoot *BSTree) {
	node, parent, pos := SearchNodeWithParent(value, treeRoot)
	if node == nil {
		return
	}
	//如果没有父节点 ，表示他就是父节点

	//没有子节点的情况 直接删除
	if node.isLeaf() {
		if parent == nil {
			*treeRoot = *(*BSTree)(nil)
			return
		}
		if pos == "left" {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
	} else if single := getSingleChild(node); single != nil { //表示要删除的节点 ，只有一个字节点
		if parent == nil {
			if node.Left != nil {
				*treeRoot = *node.Left
			} else {
				*treeRoot = *node.Right
			}
			return
		}
		if pos == "left" {
			parent.Left = single
		} else {
			parent.Right = single
		}
	} else {
		//先找右节点的最小值,就是继任者
		successor := MinNode(node.Right)
		DelNode(successor.Value, treeRoot)
		node.Value = successor.Value

	}
}

//打印空格

type BSTrees []*BSTree

func (this BSTrees) IsAllNil() bool {
	for _, v := range this {
		if v != nil {
			return false
		}
	}
	return true
}

//func PrintBSTree(trees BSTrees, maxLevel, curLevel int) {
//	if len(trees) == 0 || trees.IsAllNil() {
//		return
//	}
//
//	floor := maxLevel - curLevel
//	//左边空格数
//	leftBlanks := math.Pow(2.0, float64(floor))        //    2的 （总层高-当前层） 幂
//	rightBlanks := math.Pow(2.0, float64(floor+1)) - 1 //  2的(总层高-当前层+1)   幂
//	printBlanks(leftBlanks)                            //打印左边边空格
//
//	newNodes := make(BSTrees, 0)
//
//	for _, v := range trees {
//		if v != nil {
//			fmt.Print(v.Value)
//			newNodes = append(newNodes, v.Left, v.Right)
//		} else {
//			printBlanks(1)
//			newNodes = append(newNodes, nil, nil)
//		}
//		printBlanks(rightBlanks) //打印右边空格
//	}
//	fmt.Print("\n")
//
//	//画连接线
//	//画线
//	lineNums := math.Pow(2, float64(floor-1))
//	for i := 1.0; i <= lineNums; i++ {
//		for _, node := range trees {
//			printBlanks(leftBlanks - i) //左边线做空格
//			if node == nil {
//				printBlanks(lineNums*2 + i + 1)
//				continue
//			}
//			if node.Left != nil {
//				fmt.Print("/")
//			} else {
//				printBlanks(1)
//			}
//			printBlanks(2*i - 1) //左边线的右空格
//			if node.Right != nil {
//				fmt.Print("\\") //右斜线
//			} else {
//				printBlanks(1)
//			}
//			printBlanks(2*lineNums - i)
//		}
//		fmt.Print("\n")
//	}
//	PrintBSTree(newNodes, maxLevel, curLevel+1)
//}

func PrintBSTree(nodes BSTrees, maxlevel int, currlevel int) {
	if len(nodes) == 0 || nodes.IsAllNil() {
		return
	}
	floor := maxlevel - currlevel
	leftBlanks := math.Pow(2.0, float64(floor)) - 1    //元素左边的空格数
	rightBlanks := math.Pow(2.0, float64(floor+1)) - 1 //元素左边的空格数
	printBlanks(leftBlanks)                            //先打左边空格

	newNodes := make(BSTrees, 0)
	for _, node := range nodes {
		if node != nil {
			fmt.Print(node.Value)
			newNodes = append(newNodes, node.Left, node.Right)
		} else {
			printBlanks(1) //打印一个空格
			newNodes = append(newNodes, nil, nil)
		}
		printBlanks(rightBlanks) //打右边空格
	}
	fmt.Print("\n")

	//画线
	lineNums := math.Pow(2, float64(floor-1))
	for i := 1.0; i <= lineNums; i++ {
		for _, node := range nodes {
			printBlanks(leftBlanks - i) //左边线做空格
			if node == nil {
				printBlanks(lineNums*2 + i + 1)
				continue
			}
			if node.Left != nil {
				fmt.Print("/")
			} else {
				printBlanks(1)
			}
			printBlanks(2*i - 1) //左边线的右空格
			if node.Right != nil {
				fmt.Print("\\") //右斜线
			} else {
				printBlanks(1)
			}
			printBlanks(2*lineNums - i)
		}
		fmt.Print("\n")
	}
	PrintBSTree(newNodes, maxlevel, currlevel+1)
}
