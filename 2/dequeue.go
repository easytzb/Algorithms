//Dequeue sort.
//Explain how you would sort a deck of cards, with the restric-tion that
//the only allowed operations are to look at the values of the top two cards,
//to exchange the top two cards, and to move the top card to the bottom of the deck.
package main

import (
	"container/list"
	//"fmt"
)

func dequeue() {
	c := shuffle()
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
	//fmt.Println(c)
}
