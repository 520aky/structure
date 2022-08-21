package main

import "fmt"

//赢三张算法

func main() {
	cardBuffer := RandCardList()
	cards1 := make(Cards, 3)
	cards2 := make(Cards, 3)

	//cardBuffer = []int{28, 16, 18, 5, 18, 31}  //235吃豹子

	//fmt.Println(cardBuffer)
	for i := 0; i < 6; i++ {
		color := GetColor(cardBuffer[i])
		value := GetValue(cardBuffer[i])
		card := &Card{
			Color: color,
			Value: value,
		}
		if i > 2 {
			cards2[i-3] = card
		} else {
			cards1[i] = card
		}
	}
	fmt.Println()
	fmt.Println("=====cards1=======")
	for _, card := range cards1 {
		fmt.Printf("%s ", card.GetCardNameByCard())
	}
	fmt.Println()
	fmt.Println("=====cards2=======")
	for _, card := range cards2 {
		fmt.Printf("%s ", card.GetCardNameByCard())
	}
	fmt.Println()
	//b := GetCardType(cards1)
	//p := GetCardType(cards2)

	fmt.Println(isOvercomePrev(cards1, cards2))
}
