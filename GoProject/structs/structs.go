package main

import "fmt"

type position struct {
	x float32
	y float32
}

type badGuy struct {
	name   string
	health int
	pos    position
}

func whereIsBadGuy(b badGuy) {
	x := b.pos.x
	y := b.pos.y
	fmt.Println("[", x, y, "]")
}

func main() {

	var p position
	p.y = 5
	p.x = 4

	cor := position{10, 11}

	fmt.Println(cor.x, cor.y)
	fmt.Println(p.x, p.y)

	bb := badGuy{"Paul is a bad guy", 100, cor}
	fmt.Println(bb)
	whereIsBadGuy(bb)
}
