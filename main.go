package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Obstgarten struct {
	// cherries, apples, plums and pears
	fruits [4]int

	crow int
}

func main() {
	const N = 10000
	fmt.Println("starting...")
	rand.Seed(131)
	wons := 0
	for i := 0; i < N; i++ {
		game := NewGame()
		won := game.play()
		if won {
			wons++
		}
	}
	percent := math.Floor(0.5 + (100.0 / float64(N) * float64(wons)))
	fmt.Printf("games won: %v out of %v (%v%%)", wons, N, percent)

}

func NewGame() Obstgarten {
	o := Obstgarten{}
	for i := 0; i < len(o.fruits); i++ {
		o.fruits[i] = 10
	}
	return o
}

func (o *Obstgarten) play() bool {
	for !o.over() {
		o.step()

	}
	return o.won()
}

func (o Obstgarten) fruitsLeft() int {
	n := 0
	for i := 0; i < len(o.fruits); i++ {
		n += o.fruits[i]
	}
	return n
}

func (o Obstgarten) won() bool {
	return o.fruitsLeft() == 0
}

func (o Obstgarten) lost() bool {
	return o.crow == 9
}

func (o Obstgarten) over() bool {
	return o.lost() || o.won()
}

func (o *Obstgarten) step() {
	n := rollDice()
	switch n {
	case 0, 1, 2, 3:
		if o.fruits[n] > 0 {
			o.fruits[n]--
		}

	// basket
	case 4:
		o.pick()
		o.pick()

	// crow
	case 5:
		o.crow++
	}
}

// simple strategy: pick first available fruit

func (o *Obstgarten) pick() {
	for i := 0; i < len(o.fruits); i++ {
		if o.fruits[i] > 0 {
			o.fruits[i]--
			break
		}
	}
}

func rollDice() int {
	return rand.Intn(6)
}
