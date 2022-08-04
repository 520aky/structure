package astar

import (
	"fmt"
	"math"
	. "shujujiegou/astar/grid"
	. "shujujiegou/astar/node"
	"sort"
)

type aStar struct {
	openList  []*Node
	closeList []*Node
	grid      *Grid
	startNode *Node
	endNode   *Node
	path      []*Node
	//heuristic    func(node *Node) float64
	straightCost float64 //直线代价  Number=1.0
	diagCost     float64 //对角线代价 math.Sqrt2

}

func NewAStar() *aStar {
	return &aStar{
		openList:  make([]*Node, 0),
		closeList: make([]*Node, 0),
		//path:      make([]*Node, 0),,
		diagCost:     math.Sqrt2,
		straightCost: 1.0}
}

func (a *aStar) isOpen(node *Node) bool {
	for _, v := range a.openList {
		if v == node {
			return true
		}
	}
	return false
}

func (a *aStar) isClose(node *Node) bool {
	for _, v := range a.closeList {
		if v == node {
			return true
		}
	}
	return false
}

//曼哈顿估价法 Math.abs(node.x - _endNode.x) * _straightCost + Math.abs(node.y + _endNode.y) * _straightCost;
func (a *aStar) manhattan(node *Node) float64 {
	return math.Abs(float64(node.X-a.endNode.Y))*a.straightCost + math.Abs(float64(node.Y-a.endNode.Y))*a.straightCost
}

//几何估价法
func (a *aStar) euclidian(node *Node) float64 {
	/*
	 var dx:Number=node.x - _endNode.x;
	    var dy:Number=node.y - _endNode.y;
	    return Math.sqrt(dx * dx + dy * dy) * _straightCost;
	*/
	dx := float64(node.X - a.endNode.X)
	dy := float64(node.Y - a.endNode.Y)
	return math.Sqrt(dx*dx+dy*dy) * a.straightCost
}

////对角线估价法
func (a *aStar) diagonal(node *Node) float64 {
	/*
		var dx:Number=Math.abs(node.x - _endNode.x);
		    var dy:Number=Math.abs(node.y - _endNode.y);
		    var diag:Number=Math.min(dx, dy);
		    var straight:Number=dx + dy;
		    return _diagCost * diag + _straightCost * (straight - 2 * diag);
	*/
	dx := math.Abs(float64(node.X - a.endNode.X))
	dy := math.Abs(float64(node.Y - a.endNode.Y))
	diag := math.Min(dx, dy)
	straight := dx + dy
	return a.diagCost*diag + a.straightCost*(straight-2*diag)
}

func (a *aStar) FindPathFindPath(grid *Grid) bool {
	a.grid = grid
	a.startNode = grid.StartNode
	a.endNode = grid.EndNode
	a.startNode.G = 0
	a.startNode.H = a.diagonal(a.startNode)
	a.startNode.F = a.startNode.G + a.startNode.H
	return a.search()
}

func (a *aStar) search() bool {
	var t uint = 1
	var node *Node = a.startNode
	for node != a.endNode {
		//找出相邻8个节点的xy范围值,这里是防止是边界， 超出边界下标出错
		startX := int(math.Max(0, float64(node.X-1)))
		endX := int(math.Min(float64(a.grid.NumCols-1), float64(node.X+1)))
		startY := int(math.Max(0, float64(node.Y-1)))
		endY := int(math.Min(float64(a.grid.NumRows-1), float64(node.Y+1)))

		//循环处理所有周边节点
		for i := startX; i <= endX; i++ {
			for j := startY; j <= endY; j++ {
				testNode := a.grid.GetNode(i, j)
				if testNode == node || !testNode.Walkable ||
					!a.grid.GetNode(node.X, testNode.Y).Walkable ||
					!a.grid.GetNode(testNode.X, node.Y).Walkable {
					//如果测试节点时自身，或者不能走动时，直接进入下一循环
					continue
				}
				cost := a.straightCost
				//if (!((node.x == test.x) || (node.y == test.y)))
				//如果是对角线
				if !(node.X == testNode.X || node.Y == testNode.Y) {
					cost = a.diagCost
				}

				//计算test节点的总代价
				g := node.G + cost*testNode.CostMultiplier
				h := a.diagonal(testNode)
				f := g + h

				if a.isOpen(testNode) || a.isClose(testNode) {
					if f < testNode.F {
						testNode.F = f
						testNode.H = h
						testNode.G = g
						testNode.Parent = node
					}
				} else { //还不在列表中，需要添加到open列表中
					testNode.F = f
					testNode.H = h
					testNode.G = g
					testNode.Parent = node
					a.openList = append(a.openList, testNode)
				}

			}
		}
		a.closeList = append(a.closeList, node)
		if len(a.openList) == 0 {
			fmt.Println("没有找到最佳路径无路可走")
			return false
		}

		sort.Slice(a.openList, func(i, j int) bool {
			return a.openList[i].F < a.openList[j].F
		})

		node = a.openList[0]
		a.openList = a.openList[1:] //取出第一个元素， 因为他是代价最小的值
		fmt.Printf("第%d轮取出的最佳节点为：%s\n", t, node)
		t++
	}

	a.buildPath()
	return true
}

func (a *aStar) buildPath() {
	a.path = make([]*Node, 0)
	node := a.endNode
	a.path = append(a.path, node)

	for node != a.startNode {
		node = node.Parent
		//下面操作是往切片头部添加一个元素
		a.path = append(a.path, (*Node)(nil))
		copy(a.path[1:], a.path[:]) //todo 这里可能会报错
		a.path[0] = node
	}
}

func (a *aStar) GetPath() []*Node {
	return a.path
}
