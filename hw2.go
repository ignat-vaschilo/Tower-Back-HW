package main

import "fmt"

type Node struct {
	data int
	prev *Node
	next *Node
}

type Deck struct {
	first *Node
	last  *Node
}

func (d *Deck) AddFront(val int) {
	newNode := &Node{data: val}
	if d.first == nil {
		d.first = newNode
		d.last = newNode
	} else {
		d.first.prev = newNode
		newNode.next = d.first
		d.first = newNode
	}
}

func (d *Deck) AddBack(val int) {
	newNode := &Node{data: val}
	if d.first == nil {
		d.first = newNode
		d.last = newNode
	} else {
		d.last.next = newNode
		newNode.prev = d.last
		d.last = newNode
	}
}

func (d *Deck) PopFront() int {
	value := d.first.data
	d.first = d.first.next
	d.first.prev = nil
	return value
}

func (d *Deck) PopBack() int {
	value := d.last.data
	d.last = d.last.prev
	d.last.next = nil
	return value
}

func (d *Deck) ViewOfDeck() {
	current := d.first
	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}
	fmt.Println()
}

func (d *Deck) IsExist(val int) bool {
	current := d.first
	for current != nil {
		if current.data == val {
			return true
		}
		current = current.next
	}
	return false
}

func main() {
	deck := Deck{}
	// ViewOfDeck - чтобы увидеть дэк
	deck.AddFront(5)
	deck.AddFront(7)
	deck.AddFront(9)
	deck.AddBack(0)
	deck.ViewOfDeck()
	x := deck.PopBack()
	fmt.Println(x)
	deck.ViewOfDeck()
	y := deck.PopFront()
	fmt.Println(y)
	deck.ViewOfDeck()
	fmt.Println(deck.IsExist(0))
}
