package main

import (
	"fmt"
	astar2 "shujujiegou/astar/astar"
	"shujujiegou/astar/grid"
	"strings"
)

//A星寻路

type open struct {
	parent  *open
	pos     *Pos
	f, g, h int
}

type Pos struct {
	x    int
	y    int
	view string
}
type Maps struct {
	width  int
	height int
	zones  []*Pos
}

func (m *Maps) print() {
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			fmt.Print(" " + m.getPos(x, y).view + " ")
		}
		fmt.Println()
	}
}

func (m *Maps) getPos(x, y int) *Pos {
	if x < 0 || y < 0 || x >= m.width || y >= m.height {
		return nil
	}
	for _, pos := range m.zones {
		if pos.y == y && pos.x == x {
			return pos
		}
	}
	return nil
}

func (m *Maps) updateView(x, y int, view string) {
	p := m.getPos(x, y)
	if p != nil {
		p.view = view
	}
}

func newMaps(arr []string) *Maps {
	m := &Maps{
		width:  len(strings.Split(arr[0], " ")),
		height: len(arr),
	}

	for y := 0; y < m.height; y++ {
		viewArr := strings.Split(arr[y], " ")
		for x := 0; x < m.width; x++ {
			m.zones = append(m.zones, &Pos{
				x:    x,
				y:    y,
				view: viewArr[x],
			})
		}
	}
	return m
}

func main() {
	arr := []string{
		"O X O O X O O O",
		"O O O O X O O O",
		"X X O O X O O O",
		"O O O O O O X X",
		"O O O X O O X O",
	}

	grid := grid.NewGrid2(arr)

	//设置障碍物
	//grid := grid.NewGrid(8,5)
	//grid.GetNode(4, 0).Walkable = false
	//grid.GetNode(4, 1).Walkable = false
	//grid.GetNode(4, 2).Walkable = false
	//grid.GetNode(3, 3).Walkable = false
	//grid.GetNode(3, 4).Walkable = false

	grid.SetStartNode(0, 0)
	grid.SetEndNode(7, 4)

	astar := astar2.NewAStar()
	astar.FindPathFindPath(grid)
	for _, v := range astar.GetPath() {
		fmt.Println(v)
	}
}
