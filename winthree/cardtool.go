package main

import (
	"math/rand"
	"sort"
	"time"
)

func randInt(min, max int) int {
	if min >= max {
		return max
	}
	return rand.Intn(max-min) + min
}

func RandCardList() []int {
	count := len(CardData)
	cardBuffer := make([]int, count)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		ranOne := randInt(0, count-1-i)
		CardData[ranOne], CardData[count-1-i] = CardData[count-1-i], CardData[ranOne]
	}
	copy(cardBuffer, CardData)
	return cardBuffer
}

func GetCardType(cards Cards) int {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Value < cards[j].Value
	})

	cardType := UNDEFINE

	if isBaoZi(cards) {
		cardType = BAO_ZI
	} else if isTongHuaShun(cards) {
		cardType = TONG_HUA_SHUN
	} else if isTongHua(cards) {
		cardType = TONG_HUA
	} else if isShunZi(cards) {
		cardType = SHUN_ZI
	} else if isDuiZi(cards) {
		cardType = DUI_ZI
	}
	return cardType
}

//比牌接口函数
//--@比牌接口函数
//--@ my_Cards, 本家牌,
//--@ next_Cards,下家牌,
//--@ ret true/false

func isOvercomePrev(myCards, nextCards Cards) bool {
	myCardType := GetCardType(myCards)
	nextCardType := GetCardType(nextCards)
	if myCardType != nextCardType {
		return CardTypeDifferent(myCards, nextCards, myCardType, nextCardType)
	} else {
		return CardTypeSame(myCards, nextCards, nextCardType)
	}
}

/*
同牌型的牌比较就要分别处理了：
豹子：比较单张牌牌值

同花顺：比较第三张牌，同时考虑A23特殊顺子情况

同花：从第三张牌开始依次比较

顺子：比较第三张牌，同时考虑A23特殊顺子情况

对牌：首先比较第二张，因为第二张一定是构成对子的那张牌。若相同则再比对（第一张+第三张）

另外：赢三张规定，三张牌值完全相同的情况下，比牌者输
————————————————
版权声明：本文为CSDN博主「九日王朝」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/sm9sun/article/details/72466311
*/

//当牌型相同是处理流程
func CardTypeSame(myCards, nextCards Cards, cardType int) bool {
	//豹子  三张牌一样比较牌点大小, 同点本家输
	if cardType == BAO_ZI {
		return myCards[0].Value-nextCards[0].Value > 0
	}

	//同花顺  判断第一张牌的大小 同时考虑最小同花顺A23
	if cardType == TONG_HUA_SHUN {
		myIsA23 := isA23(myCards)
		nextIsA23 := isA23(myCards)
		//两个都是A23,本家输
		if myIsA23 && nextIsA23 {
			return false
		} else if myIsA23 || nextIsA23 {
			//其中有一个是A23 A23的输
			return !myIsA23
		} else {
			//两个都不是A23 ，判断第一张牌牌的大小,同点本家输
			return myCards[0].Value-nextCards[0].Value > 0
		}
	}

	//同花
	if cardType == TONG_HUA {
		//依次从最大牌开始比牌点大小
		if myCards[2].Value-nextCards[2].Value > 0 {
			return true
		} else if myCards[2].Value-nextCards[2].Value < 0 {
			return false
		} else if myCards[1].Value-nextCards[1].Value > 0 {
			return true
		} else if myCards[1].Value-nextCards[1].Value < 0 {
			return false
		} else if myCards[0].Value-nextCards[0].Value > 0 {
			return true
		} else if myCards[0].Value-nextCards[0].Value < 0 {
			return false
		} else {
			//3张牌都相同点数，本家输
			return false
		}
	}
	//顺子  考虑A23最小顺子情况
	if cardType == SHUN_ZI {
		myIsA23 := isA23(myCards)
		nextIsA23 := isA23(nextCards)
		//两个都是A23的情况 ,
		if myIsA23 && nextIsA23 {
			return false
		} else if myIsA23 || nextIsA23 {
			return !myIsA23
		} else {
			return myCards[0].Value-nextCards[0].Value > 0
		}
	}
	//对子 , 第二张牌肯定是组成对子的牌
	if cardType == DUI_ZI {
		tmp := myCards[1].Value - nextCards[1].Value
		if tmp > 0 {
			return true
		} else if tmp < 0 {
			return false
		} else {
			//第二张牌相同点数, 说明对子相同 ，比较剩余一张牌点数大小，
			//因为不能判断 和第几张是对子， 所以相加剩余两个比较点数大小，相同的话本家输
			return (myCards[0].Value+myCards[2].Value)-(nextCards[0].Value+nextCards[2].Value) > 0
		}
	}

	//单牌处理  --依次从最大牌开始比牌点大小
	if myCards[2].Value-nextCards[2].Value > 0 {
		return true
	} else if myCards[2].Value-nextCards[2].Value < 0 {
		return false
	} else if myCards[1].Value-nextCards[1].Value > 0 {
		return true
	} else if myCards[1].Value-nextCards[1].Value < 0 {
		return false
	} else if myCards[0].Value-nextCards[0].Value > 0 {
		return true
	} else if myCards[0].Value-nextCards[0].Value < 0 {
		return false
	} else {
		//3张牌都相同点数，本家输
		return false
	}
}

//当牌型不同时的处理
func CardTypeDifferent(myCards, nextCards Cards, myCardType, nextCardType int) bool {
	myCardBaoZi := false
	nextCardBaoZi := false
	if myCardType == BAO_ZI {
		myCardBaoZi = true
	}
	if nextCardType == BAO_ZI {
		nextCardBaoZi = true
	}

	////如果没有两个都没有豹子时， 直接比两个cardType的大小  myCardType -  nextCardType > 0  表示本家大于下家
	if myCardBaoZi == nextCardBaoZi { //因为这个函数时处理不不同的， 说明两个类型不可能是同时为豹子，相等则说明都不是豹子
		return myCardType-nextCardType > 0 //他们两个是不同的，所以不可能出现等于0的情况
	}

	//这里不可能出现两个都是豹子的情况， 下面流程只能出现有一家为豹子的情况
	//当有一个为豹子时，判断对家是否是235并且不能是同花(235管豹子)
	//if myCardBaoZi || nextCardBaoZi {
	//判断是否是235
	my235 := is235(myCards)
	next235 := is235(nextCards)
	if my235 { //如果本家为235说明下家为豹子
		return !isTongHua(myCards) //如果为同花，则不能压豹子
	} else if next235 { //如果下家为235说明本家为豹子,下家是同花本家就赢
		return isTongHua(nextCards)
	} else {
		//如果两家都不是235， 本家为豹子就本家赢， 因为这个流程是 有且只有一家为豹子
		return myCardBaoZi
	}
	//}

}
