package main

import (
	. "shujujiegou/tree/btree"
)

/*
         10
    8           12
  7   13    11     14
 5         22     33
4                32
*/

/**/

func main2() {
	root := NewBTree(10)
	root.ConnectLeft(8).ConnectRight(12)
	{
		root.Left.ConnectLeft(7).ConnectRight(13)
		root.Left.Left.ConnectRight(5)
		root.Left.Left.Right.ConnectRight(4)
		root.Right.ConnectLeft(11).ConnectRight(14)
		root.Right.Left.ConnectLeft(22)
		root.Right.Right.ConnectLeft(33)
		root.Right.Right.Left.ConnectLeft(32)
	}

	//maxlevel := root.GetLevel()
	//PrintBTree([]*BTree{root}, maxlevel, 1)
	root.Postorder()

}
