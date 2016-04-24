//Dequeue sort.
//Explain how you would sort a deck of cards, with the restric-tion that
//the only allowed operations are to look at the values of the top two cards,
//to exchange the top two cards, and to move the top card to the bottom of the deck.
package main

import (
	"container/list"
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
	} else {
		return 0
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

func dequeue(c []Card) []Card {
	len := len(c)
	l := list.New()
	for i := 0; i < len; i++ {
		l.PushBack(c[i])
	}

	for i := 0; i < len; i++ {
		//find smallest
		for j := len - 1 - i; j > 0; j-- {
			if l.Front().Value.(Card).less(l.Front().Next().Value.(Card)) == 1 {
				l.MoveToBack(l.Front().Next())
			} else {
				l.MoveToBack(l.Front())
			}
		}

		for j := i; j > 0; j-- {
			l.MoveToBack(l.Front().Next())
		}
		l.MoveToBack(l.Front())
	}

	index := 0
	for e := l.Front(); e != nil; e = e.Next() {
		c[index] = e.Value.(Card)
		index++
	}
	return c
}

// Can be run as follow:
// go run deck.go
func main() {
	fmt.Println(dequeue(getDeck()))
}
