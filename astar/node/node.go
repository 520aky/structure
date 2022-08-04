package node

import "fmt"

//参考：http://www.manongjc.com/detail/50-xkxkwyirjaqbvag.html

type Node struct {
	X        int
	Y        int
	F        float64
	G        float64
	H        float64
	Walkable bool //是否可穿越（通常把障碍物节点设置为false）
	Parent   *Node
	//这里就是一个1.0
	CostMultiplier float64 //Number=1.0;//代价因子
}

func NewNode(x, y int, b bool) *Node {
	return &Node{
		X:              x,
		Y:              y,
		CostMultiplier: 1.0,
		Walkable:       b,
	}
}

/*
public function toString():String{
    var fmr:NumberFormat = new NumberFormat();
    fmr.mask = "#.0";
    return "x=" + this.x.toString() + ",y=" + this.y.toString() + ",g=" + fmr.format(this.g) + ",h=" + fmr.format(this.h) + ",f=" + fmr.format(this.f);
}

*/
func (n *Node) String() string {
	return fmt.Sprintf("x=%d,y=%d,g=%f,h=%f,f=%f", n.X, n.Y, n.G, n.H, n.F)
}
