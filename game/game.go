package main

import "fmt"

func main() {
	var i1 Item
	fmt.Println(i1)
	fmt.Printf("i1: %#v\n", i1)

	i2 := Item{1, 2}
	fmt.Printf("i2: %#v\n", i2)

	i3 := Item{
		X: 12,
		Y: 13,
	}
	fmt.Printf("i3: %#v\n", i3)

	i2.Move(10, 20)
	fmt.Printf("i2: %#v\n", i2)

	p1 := Player{
		Item: Item{1, 2},
		Name: "Player 1",
	}
	fmt.Printf("p1: %#v\n", p1)
	fmt.Printf("p1.X: %#v\n", p1.X)

	ms := []mover{
		&i1,
		&p1,
		&i2,
	}

	moveAll(ms, 10, 20)
	for _, m := range ms {
		fmt.Printf("m: %#v\n", m)
	}
}

type mover interface {
	Move(x, y int)
}

func moveAll(ms []mover, x, y int) {
	for _, m := range ms {
		m.Move(x, y)
	}
	k := Jade

	fmt.Println("K:", k)
}

const (
	Jade Key = iota + 1
	Copper
	Crystal
)

func (k Key) String() string {
	switch k {
	case Jade:
		return "Jade"
	case Copper:
		return "Cooper"
	case Crystal:
		return "crystal"
	}

	return fmt.Sprintf("<Key %d", k)

}

type Key byte

// func NewItem(x, y int) Item{}
// func NewItem(x, y int) *Item{}
/** func NewItem(x, y int) (Item, error){
	if x < 0 || x > maxX {
		return Item{}, fmt.Errorf("x is out of range")
	}
	if y < 0 || y > maxY {
		return Item{}, fmt.Errorf("y is out of range")
	}
	return Item{x, y}, nil
}
*/

type Player struct {
	Item
	Name string
}

// i is called "receiver"
func (i *Item) Move(x, y int) {
	i.X = x
	i.Y = y
}

func NewItem(x, y int) (*Item, error) {
	if x < 0 || x > maxX {
		return nil, fmt.Errorf("x is out of range")
	}
	if y < 0 || y > maxY {
		return nil, fmt.Errorf("y is out of range")
	}
	return &Item{x, y}, nil
}

const (
	maxX = 100
	maxY = 100
)

type Item struct {
	X int
	Y int
}
