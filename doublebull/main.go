package main

import (
	"fmt"
)

func main() {

	cardBuffer := RandCardList()
	fmt.Println(CardData)

	//cardBuffer = []uint8{21, 36, 43, 47, 32}

	cards1 := make(Cards, 5)
	cards2 := make(Cards, 5)
	for i := 0; i < 10; i++ {
		color := GetColor(cardBuffer[i])
		value := GetValue(cardBuffer[i])
		count := GetCount(value)
		code := cardBuffer[i]
		card := &Card{
			Color: color,
			Value: value,
			Count: count,
			Code:  int(code),
		}
		if i >= 5 {
			cards2[i-5] = card
		} else {
			cards1[i] = card
		}
	}
	for _, card := range cards1 {
		fmt.Println(card.GetCardNameByCard())
	}

	for _, card := range cards2 {
		fmt.Println(card.GetCardNameByCard())
	}

	//sort.Sort(cards1)
	//sort.Sort(cards2)
	fmt.Println(cards1.getTypeByCards())
	fmt.Println(cards2.getTypeByCards())
	fmt.Println(bankerIsWin(cards1, cards2))
}
