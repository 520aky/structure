package grid

import (
	"shujujiegou/astar/node"
	"strings"
)

/*
	    private var _startNode:Node;//开始节点
		private var _endNode:Node;//目标节点
		private var _nodes:Array;//节点数组
		private var _numCols:int;//列数
		private var _numRows:int;//行数
*/

type Grid struct {
	StartNode *node.Node
	EndNode   *node.Node
	Nodes     [][]*node.Node
	NumCols   int
	NumRows   int
}

func NewGrid(numCols, numRows int) *Grid {
	grid := &Grid{
		NumCols: numCols,
		NumRows: numRows,
	}
	grid.Nodes = make([][]*node.Node, numCols)
	for i := 0; i < numCols; i++ {
		grid.Nodes[i] = make([]*node.Node, numRows)
	}

	for i := 0; i < numCols; i++ {
		for j := 0; j < numRows; j++ {
			grid.Nodes[i][j] = node.NewNode(i, j, true)
		}
	}
	return grid
}

//这里是输出一个字符串数组，来初始化
func NewGrid2(arr []string) *Grid {
	grid := &Grid{
		NumCols: len(strings.Split(arr[0], " ")),
		NumRows: len(arr),
	}

	//width:  len(strings.Split(arr[0], " ")),
	//height: len(arr),

	grid.Nodes = make([][]*node.Node, grid.NumCols)
	for i := 0; i < grid.NumCols; i++ {
		grid.Nodes[i] = make([]*node.Node, grid.NumRows)
	}

	for y := 0; y < grid.NumRows; y++ {
		viewArr := strings.Split(arr[y], " ")
		for x := 0; x < grid.NumCols; x++ {
			tmp := true
			if viewArr[x] == "X" {
				tmp = false
			}
			grid.Nodes[x][y] = node.NewNode(x, y, tmp)
		}
	}
	return grid
}

func (g *Grid) GetNode(x, y int) *node.Node {
	return g.Nodes[x][y]
}

func (g *Grid) SetStartNode(x, y int) {
	g.StartNode = g.Nodes[x][y]
}

func (g *Grid) SetEndNode(x, y int) {
	g.EndNode = g.Nodes[x][y]
}

func (g *Grid) SetWalkable(x, y int, value bool) {
	g.Nodes[x][y].Walkable = value
}

func (g *Grid) GetEndNode() *node.Node {
	return g.EndNode
}

func (g *Grid) GetStartNode() *node.Node {
	return g.StartNode
}

func (g *Grid) GetNumCols() int {
	return g.NumCols
}

func (g *Grid) GetNumRows() int {
	return g.NumRows
}
