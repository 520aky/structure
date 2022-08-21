package main

import (
	"fmt"
	"math"
	"strings"
)

const (
	UNDEFINE      = 0 //        --单牌
	DUI_ZI        = 1 //          --对子
	SHUN_ZI       = 2 //         --顺子
	TONG_HUA      = 3 //   --同花
	TONG_HUA_SHUN = 4 // --同花顺
	BAO_ZI        = 5 //       --豹子

)

var CardData = []int{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
	14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26,
	27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
	40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52,
}

type Card struct {
	Color int
	Value int
}

func GetColor(val int) int {
	return int(math.Ceil(float64(val) / float64(13)))
}

func GetValue(val int) int {
	t := int(val % 13)
	if t == 0 {
		return 13
	}
	if t == 1 {
		return 14 //如果是 A 则标记为14， 因为A是赢三张里面最大的单点牌
	}
	return t
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
	case 14:
		sb.WriteString("A")
	case 13:
		sb.WriteString("K")
	case 12:
		sb.WriteString("Q")
	case 11:
		sb.WriteString("J")
	default:
		sb.WriteString(fmt.Sprintf("%d", this.Value))
	}

	return sb.String()
}

type Cards []*Card

//豹子
func isBaoZi(cards Cards) bool {
	return cards[0].Value == cards[2].Value
}

//同花
func isTongHua(cards Cards) bool {
	if cards[0].Color == cards[1].Color && cards[0].Color == cards[2].Color {
		return true
	}
	return false
}

//顺子
func isShunZi(cards Cards) bool {
	if isA23(cards) {
		return true
	}
	if cards[2].Value-cards[1].Value == 1 && cards[1].Value-cards[0].Value == 1 {
		return true
	}
	return false
}

//同花顺
func isTongHuaShun(cards Cards) bool {
	if isTongHua(cards) && isShunZi(cards) {
		return true
	}
	return false
}

//对子 两张牌牌值相等，但第一张与第三张不能相等，否则就是豹子了
func isDuiZi(cards Cards) bool {
	if cards[0].Value != cards[2].Value {
		if cards[0].Value == cards[1].Value {
			return true
		}
		if cards[1].Value == cards[2].Value {
			return true
		}
	}
	return false
}

//是否是 A23 最小的顺子
func isA23(cards Cards) bool {
	if cards[0].Value == 2 && cards[1].Value == 3 && cards[2].Value == 14 {
		return true
	}
	return false
}

//是否是235
func is235(cards Cards) bool {
	if cards[0].Value == 2 && cards[1].Value == 3 && cards[2].Value == 5 {
		return true
	}
	return false
}
