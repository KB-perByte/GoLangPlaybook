package main

import (
	"fmt"
)

func addOne(x *int) {
	*x = *x + 1
}

type position struct {
	x float32
	y float32
}

type badGuy struct {
	name   string
	health int
	pos    position
}

func whereIsBadGuy(b *badGuy) {
	x := b.pos.x
	y := b.pos.y
	fmt.Println("[", x, y, "]")
}

func main() {
	x := 5
	var y int = 10
	fmt.Println(x)

	cor := position{10, 11}

	xPtr := &x
	fmt.Println(xPtr)

	var yPtr *int = &y
	fmt.Println(yPtr)

	addOne(xPtr)
	fmt.Println(x)

	bb := badGuy{"Paul is a bad guy", 100, cor}
	fmt.Println(bb)
	whereIsBadGuy(&bb)
}
