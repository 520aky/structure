package grid

import (
	"fmt"
	"testing"
)

func TestNewGrid(t *testing.T) {
	grid := NewGrid(5, 6)
	for i := 0; i < 5; i++ {
		for j := 0; j < 6; j++ {
			fmt.Print(grid.Nodes[i][j].CostMultiplier)
		}
		fmt.Println()
	}

}
