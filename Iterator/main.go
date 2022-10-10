package main

import "fmt"

type Person struct {
	FirstName, MiddleName, LastName string
}

func (p Person) NamesGenerator() <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		ch <- p.FirstName
		if p.MiddleName != "" {
			ch <- p.MiddleName
		}
		ch <- p.LastName
	}()
	return ch
}

type PersonNameIterator struct {
	person *Person
	index  int
}

func NewNamesIterator(p *Person) *PersonNameIterator {
	return &PersonNameIterator{person: p, index: -1}
}

func (p *PersonNameIterator) Next() bool {
	p.index++
	return p.index < 3
}

func (p *PersonNameIterator) Value() string {
	switch p.index {
	case 0:
		return p.person.FirstName
	case 1:
		return p.person.MiddleName
	case 2:
		return p.person.LastName
	}
	panic("unreachable")
}

type Node struct {
	Value               int
	left, right, parent *Node
}

func NewTerminalNode(v int) *Node {
	return &Node{Value: v}
}

func NewBinaryNode(value int, left, right *Node) *Node {
	n := &Node{Value: value, left: left, right: right}
	left.parent = n
	right.parent = n
	return n
}

type InOrderIterator struct {
	current, root *Node
	returnedStart bool
}

func NewInOrderIterator(root *Node) *InOrderIterator {
	i := &InOrderIterator{root: root, current: root}
	for i.current.left != nil {
		i.current = i.current.left
	}
	return i
}

func (i *InOrderIterator) Next() bool {
	if !i.returnedStart {
		i.returnedStart = true
		return true
	}

	if i.current.right != nil {
		i.current = i.current.right
		for i.current.left != nil {
			i.current = i.current.left
		}
		return true
	}

	p := i.current.parent
	for p != nil && i.current == p.right {
		i.current = p
		p = p.parent
	}
	i.current = p
	return i.current != nil
}

func (i *InOrderIterator) Value() int {
	return i.current.Value
}

func main() {
	for it := NewNamesIterator(&Person{"John", "Paul", "Doe"}); it.Next(); {
		fmt.Println(it.Value())
	}

	root := NewBinaryNode(1,
		NewBinaryNode(2,
			NewTerminalNode(3),
			NewTerminalNode(4)),
		NewBinaryNode(5,
			NewTerminalNode(6),
			NewTerminalNode(7)),
	)

	for it := NewInOrderIterator(root); it.Next(); {
		fmt.Println(it.Value())
	}

}
