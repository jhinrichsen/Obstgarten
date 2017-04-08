package main

import (
	"fmt"
)

type Obstgarten struct {
	// cherries, apples, plums and pears
	fruits [4]int

	crow int
}

func NewGame() Obstgarten {
	o := Obstgarten{}
	for i := 0; i < len(o.fruits); i++ {
		o.fruits[i] = 10
	}
	return o
}

func (Obstgarten) play() bool {
	for !over() {
		step()
	}
	return won()
}

func (Obstgarten) fruitsLeft() int {
	n := 0
	for i := 0; i < len(fruits); i++ {
		n += fruits[i]
	}
	return n
}

func (Obstgarten) won() bool {
	return fruitsLeft() == 0

}

func main() {
	fmt.Println("starting...")
	game := NewGame()
	won := game.play()
	fmt.Printf("game won: %v", won)
}
