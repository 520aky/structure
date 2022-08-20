package main

import (
	"fmt"
	"math"
	"strings"
)

/*
	A    2   3   4   5   6   7   8   9  10  J   Q   K
	1,   2,  3,  4,  5,  6,  7,  8,  9, 10, 11, 12, 13,   方片
	14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26,   梅花
	27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39,   红桃
	40, 41, 42, 43, 44, 45, 46, 47, 48, 39, 50, 51, 52,   黑桃

*/

var CardData = []uint8{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
	14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26,
	27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
	40, 41, 42, 43, 44, 45, 46, 47, 48, 39, 50, 51, 52,
}

const (
	Spade = 4
	Heart = 3
	Plum  = 2
	Block = 1
)

type Card struct {
	Color int
	Value int
	Count int
	Code  int
}

func (this *Card) String() string {
	return fmt.Sprintf("color=%d, value=%d, count=%d, code=%d", this.Color, this.Value, this.Count, this.Code)
}

func (this *Card) GetCardNameByCard() string {
	var sb strings.Builder
	if this.Color == 4 {
		sb.WriteString("黑桃")
	} else if this.Color == 3 {
		sb.WriteString("红桃")
	} else if this.Color == 2 {
		sb.WriteString("梅花")
	} else if this.Color == 1 {
		sb.WriteString("方片")
	}

	switch this.Value {
	case 13:
		sb.WriteString("K")
	case 12:
		sb.WriteString("Q")
	case 11:
		sb.WriteString("J")
	default:
		sb.WriteString(fmt.Sprintf("%d", this.Value))
	}

	sb.WriteString(fmt.Sprintf(" code = %d", this.Code))
	return sb.String()
}

func GetColor(val uint8) int {
	return int(math.Ceil(float64(val) / float64(13)))
}

func GetValue(val uint8) int {
	t := int(val % 13)
	if t == 0 {
		return 13
	}
	return t
}

func GetCount(val int) int {
	if val > 10 {
		return 10
	} else {
		return val
	}
}
