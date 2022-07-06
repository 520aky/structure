package btree

import (
	"fmt"
	"math"
)

//二叉树
type BTree struct {
	Value int
	Left  *BTree
	Right *BTree
}

func NewBTree(value int) *BTree {
	return &BTree{Value: value}
}

type BTrees []*BTree

func NewBTrees(values ...int) []*BTree {
	btress := make([]*BTree, len(values))
	for i, v := range values {
		btress[i] = NewBTree(v)
	}
	return btress
}

func (this BTrees) IsAllNil() bool {
	for _, v := range this {
		if v != nil {
			return false
		}
	}
	return true
}

func (this *BTree) String() {
	fmt.Printf("二叉树:值是%d\n", this.Value)
	if this.Left != nil {
		fmt.Printf("左节点:%v\n", this.Left)
	}
	if this.Right != nil {
		fmt.Printf("右节点:%v\n", this.Right)
	}
}

//先序遍历
func (this *BTree) Preorder() {
	if this == nil {
		return
	}
	fmt.Printf("%d->", this.Value)
	this.Left.Preorder()
	this.Right.Preorder()
}

//中序遍历
func (this *BTree) Inorder() {
	if this == nil {
		return
	}

	this.Left.Inorder()
	fmt.Printf("%d->", this.Value)
	this.Right.Inorder()
}

//后序遍历
func (this *BTree) Postorder() {
	if this == nil {
		return
	}

	this.Left.Postorder()
	this.Right.Postorder()
	fmt.Printf("%d->", this.Value)
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

//获取树的高度
func (this *BTree) GetLevel() int {
	if this == nil {
		return 0
	}

	return max(this.Left.GetLevel(), this.Right.GetLevel()) + 1
}

func (this *BTree) ConnectLeft(treeOrValue interface{}) *BTree {
	if bt, ok := treeOrValue.(*BTree); ok {
		this.Left = bt
	} else if v, ok := treeOrValue.(int); ok {
		this.Left = NewBTree(v)
	}
	return this
}

func (this *BTree) ConnectRight(treeOrValue interface{}) *BTree {
	if bt, ok := treeOrValue.(*BTree); ok {
		this.Right = bt
	} else if v, ok := treeOrValue.(int); ok {
		this.Right = NewBTree(v)
	}
	return this
}

//打印空格
func printBlanks(count float64) {
	for i := 0; i < int(count); i++ {
		fmt.Print(" ")
	}
}

/*
         10
    8           12
  7   13    11     14
          22     33
                32
*/

/*
----10---------
-8----12----
7-13-

*/

func PrintBTree(trees BTrees, maxLevel, curLevel int) {
	if len(trees) == 0 || trees.IsAllNil() {
		return
	}

	floor := maxLevel - curLevel
	//左边空格数
	leftBlanks := math.Pow(2.0, float64(floor))        //    2的 （总层高-当前层） 幂
	rightBlanks := math.Pow(2.0, float64(floor+1)) - 1 //  2的(总层高-当前层+1)   幂
	printBlanks(leftBlanks)                            //打印左边边空格

	newNodes := make(BTrees, 0)

	for _, v := range trees {
		if v != nil {
			fmt.Print(v.Value)
			newNodes = append(newNodes, v.Left, v.Right)
		} else {
			printBlanks(1)
			newNodes = append(newNodes, nil, nil)
		}
		printBlanks(rightBlanks) //打印右边空格
	}
	fmt.Print("\n")

	//画连接线
	//画线
	lineNums := math.Pow(2, float64(floor-1))
	for i := 1.0; i <= lineNums; i++ {
		for _, node := range trees {
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
	PrintBTree(newNodes, maxLevel, curLevel+1)
}
