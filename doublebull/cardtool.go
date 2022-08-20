package main

import (
	"math/rand"
	"sort"
	"time"
)

//洗牌

func randInt(min, max int) int {

	if min >= max {
		return max
	}
	return rand.Intn(max-min) + min
}

/*
CardType =
{

    NOT_NIU=0,        --没牛
    NIU_1 =1,         --牛一
    NIU_2 =2,         --牛二
    NIU_3 =3,         --牛三
    NIU_4 =4,         --牛四
    NIU_5 =5,         --牛五
    NIU_6 =6,         --牛六
    NIU_7 =7,         --牛七
    NIU_8 =8,         --牛八
    NIU_9 =9,         --牛九
    NIU_NIU =10,      --牛牛
    SILVER_NIU =11,   --银牛
    GOLD_NIU=12,      --金牛
    BOMB = 13,        --炸弹
    SMALL_NIU = 14,   --五小牛
}
*/

const (
	NOT_NIU    = 0
	NIU_1      = 1
	NIU_2      = 2
	NIU_3      = 3
	NIU_4      = 4
	NIU_5      = 5
	NIU_6      = 6
	NIU_7      = 7
	NIU_8      = 8
	NIU_9      = 9
	NIU_NIU    = 10
	SILVER_NIU = 11
	GOLD_NIU   = 12
	BOMB       = 13
	SMALL_NIU  = 14
)

func RandCardList() []uint8 {
	var count = len(CardData)
	cardBuffer := make([]uint8, count)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		ranOne := randInt(0, count-1-i)
		CardData[ranOne], CardData[count-1-i] = CardData[count-1-i], CardData[ranOne]
	}
	copy(cardBuffer, CardData)
	return cardBuffer

}

type Cards []*Card

func (this Cards) Less(i, j int) bool {
	if this[i].Value < this[j].Value {
		return true
	} else if this[i].Value > this[j].Value {
		return false
	} else {
		return this[i].Color < this[j].Color
	}
}
func (this Cards) Len() int {
	return len(this)
}

func (this Cards) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

/*
银牛：1张10 加4张大于10的牌

金牛：5张大于10的牌

炸弹：存在四张相同的牌

五小牛：五张牌总数值相加小于等于10
*/

//五小牛
func (this Cards) isSmallNiu() bool {
	sum := 0
	for _, card := range this {
		sum += card.Count
	}

	if sum <= 10 {
		return true
	} else {
		return false
	}
}

//炸弹
func (this Cards) isBomb() bool {
	if this[0].Value == this[3].Value {
		return true
	} else if this[1].Value == this[4].Value {
		return true
	}
	return false
}

//金牛
func (this Cards) isGoldNiu() bool {
	if this[0].Value > 10 {
		return true
	}
	return false
}

//银牛
func (this Cards) isSilverNiu() bool {
	if this[0].Value == 10 && this[1].Value > 10 {
		return true
	}
	return false
}

//其他牛
func (this Cards) getNiuByCards() int {
	lave := 0
	for _, card := range this {
		lave += card.Count
	}

	lave = lave % 10
	for i := 0; i < len(this)-1; i++ {
		for j := i + 1; j < len(this); j++ {
			if (this[i].Count+this[j].Count)%10 == lave {
				if lave == 0 {
					return 10
				} else {
					return lave
				}
			}
		}
	}
	return 0
}

func (this Cards) getTypeByCards() int {
	sort.Sort(this)
	cardType := 0
	if this.isSmallNiu() {
		cardType = SMALL_NIU
	} else if this.isBomb() {
		cardType = BOMB
	} else if this.isGoldNiu() {
		cardType = GOLD_NIU
	} else if this.isSilverNiu() {
		cardType = SILVER_NIU
	} else {
		cardType = this.getNiuByCards()
	}

	return cardType
}

//banker_Cards  庄家
//other_Cards  闲家
//return 庄家是否赢
func bankerIsWin(bCards Cards, pCards Cards) bool {
	bankCardType := bCards.getTypeByCards()
	playerCardType := pCards.getTypeByCards()
	
	if bankCardType != playerCardType {
		return bankCardType > playerCardType
	}

	if bankCardType == SMALL_NIU {
		return true
	}

	if bankCardType == BOMB {
		return bCards[2].Value > pCards[2].Value
	}

	if bankCardType == GOLD_NIU {
		return compByCardsValue(*pCards[4], *bCards[4])
	}

	if bankCardType == SILVER_NIU {
		return compByCardsValue(*pCards[4], *bCards[4])
	}

	if bankCardType == NIU_NIU {
		return compByCardsValue(*pCards[4], *bCards[4])
	}
	if bankCardType == NOT_NIU {
		return compByCardsValue(*pCards[4], *bCards[4])
	}

	return true
}

func compByCardsValue(b, p Card) bool {
	if b.Value < p.Value {
		return true
	}
	if b.Value > p.Value {
		return false
	}

	return b.Color < p.Color
}
