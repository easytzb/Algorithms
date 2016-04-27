//Deck sort.
//Explain how you would put a deck of cards in order by suit
//(in the order spades, hearts, clubs, diamonds) and
//by rank within each suit, with the restriction that the cards
//must be laid out face down in a row, and the only allowed
//operations are to check the values of two cards and
//to exchange two cards (keeping them face down).
package main

//import "fmt"

func deck() {
	c := shuffle()
	len := len(c)
	for step := len >> 1; step > 0; step >>= 1 {
		for i := step; i < len; i++ {
			tmp := c[i]
			min_val_index := i
			for j := i - step; j >= 0 && tmp.less(c[j]) == 1; j -= step {
				c[j+step] = c[j]
				min_val_index = j
			}

			if i != min_val_index {
				c[min_val_index] = tmp
			}
		}
	}

	//fmt.Println(c)
}
