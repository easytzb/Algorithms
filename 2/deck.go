//Deck sort.
//Explain how you would put a deck of cards in order by suit
//(in the order spades, hearts, clubs, diamonds) and
//by rank within each suit, with the restriction that the cards
//must be laid out face down in a row, and the only allowed
//operations are to check the values of two cards and
//to exchange two cards (keeping them face down).
package main

import (
	"fmt"
	"math/rand"
)

type Card struct {
	//suit of cards. 4 spades, 3 hearts, 2 clubs, 1 diamonds
	suit int

	//values of cards 0-13
	val int
}

//牌面大小的比较
func (c1 Card) less(c2 Card) int {
	v1 := c1.suit*100 + c1.val
	v2 := c2.suit*100 + c2.val
	if v1 < v2 {
		return 1
	} else if v1 == v2 {
		return 0
	} else {
		return -1
	}
}

//洗牌
func getDeck() []Card {
	c := make([]Card, 52)
	t := make([]Card, 52)
	index := 0

	for i := 1; i < 5; i++ {
		for j := 1; j < 14; j++ {
			t[index] = Card{i, j}
			index++
		}
	}

	for i := 0; i < 52; i++ {
		j := rand.Intn(i + 1)
		c[i] = c[j]
		c[j] = t[i]
	}

	return c
}

func sortDeck(c []Card) []Card {
	len := len(c)
	for step := len >> 1; step > 0; step = step >> 1 {
		for i := step; i < len; i++ {
			tmp := c[i]
			tmp_index := i
			for j := i - step; j >= 0 && tmp.less(c[j]) == 1; j -= step {
				c[j+step] = c[j]
				tmp_index = j
			}
			if i != tmp_index {
				c[tmp_index] = tmp
			}
		}
	}
	return c
}

// Can be run as follow:
// go run deck.go
func main() {
	fmt.Println(sortDeck(getDeck()))
}
