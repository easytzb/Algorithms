package main

import (
	"math/rand"
)

type Card struct {
	//suit of cards. 4 spades, 3 hearts, 2 clubs, 1 diamonds
	suit int

	//values of cards 0-13
	val int
}

func (c1 Card) less(c2 Card) int {
	v1 := c1.suit*100 + c1.val
	v2 := c2.suit*100 + c2.val
	if v1 < v2 {
		return 1
	} else {
		return 0
	}
}

func shuffle() []Card {
	c := make([]Card, 52)
	t := make([]Card, 52)
	index := 0

	for i := 1; i < 5; i++ {
		for j := 1; j < 14; j++ {
			t[index] = Card{i, j}
			index++
		}
	}

	//just like bucket sort
	for i := 0; i < 52; i++ {
		//为了达到洗牌的随机效果，将之前已入桶的牌j放到当前i号位置
		j := rand.Intn(i + 1)
		c[i] = c[j]
		c[j] = t[i]
	}

	return c
}
