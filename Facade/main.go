package main

import "fmt"

type Buffer struct {
	width, height int
	buffer        []rune
}

func NewBuffer(width, height int) *Buffer {
	return &Buffer{width, height, make([]rune, width*height)}
}

func (b *Buffer) At(x, y int) rune {
	return b.buffer[b.width*y+x]
}

func (b *Buffer) Set(x, y int, r rune) {
	b.buffer[b.width*y+x] = r
}

func (b *Buffer) String() string {
	var s string
	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			s += string(b.At(x, y))
		}
		s += ""
	}
	return s
}

type ViewPort struct {
	buffer *Buffer
	offset int
}

func NewViewPort(buffer *Buffer, offset int) *ViewPort {
	return &ViewPort{buffer, offset}
}

func (v *ViewPort) GetCharacterAt(index int) rune {
	return v.buffer.At(index+v.offset, 0)
}

type Console struct {
	buffer    []*Buffer
	viewports []*ViewPort
	offset    int
}

func NewConsole() *Console {
	b := NewBuffer(100, 100)
	v := NewViewPort(b, 0)
	return &Console{[]*Buffer{b}, []*ViewPort{v}, 0}
}

func (c *Console) GetCharacterAt(index int) rune {
	return c.viewports[0].GetCharacterAt(index)
}

func main() {
	c := NewConsole()
	c.buffer[0].Set(0, 0, 'a')
	fmt.Println(c.GetCharacterAt(0))
}
