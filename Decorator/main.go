package main

import "fmt"

type Bird struct {
	Age int
}

func (b *Bird) Fly() {
	if b.Age >= 10 {
		println("flying")
	}
}

type Lizard struct {
	Age int
}

func (l *Lizard) Crawl() {
	if l.Age < 10 {
		println("crawling")
	}
}

type Dragon struct {
	Bird
	Lizard
}

type Aged interface {
	Age() int
	SetAge(age int)
}

type Bird2 struct {
	age int
}

func (b *Bird2) Age() int {
	return b.age
}

func (b *Bird2) SetAge(age int) {
	b.age = age
}

func (b *Bird2) Fly() {
	if b.age >= 10 {
		println("flying")
	} else {
		println("too young to fly")
	}
}

type Lizard2 struct {
	age int
}

func (d *Lizard2) Age() int {
	return d.age
}

func (d *Lizard2) SetAge(age int) {
	d.age = age
}

func (d *Lizard2) Crawl() {
	if d.age < 10 {
		println("crawling")
	} else {
		println("too old to crawl")
	}
}

type Dragon2 struct {
	bird   Bird2
	lizard Lizard2
}

func (d *Dragon2) Age() int {
	return d.bird.age
}

func (d *Dragon2) SetAge(age int) {
	d.bird.age = age
	d.lizard.age = age
}

func (d *Dragon2) Fly() {
	d.bird.Fly()
}

func (d *Dragon2) Crawl() {
	d.lizard.Crawl()
}

func NewDragon2() *Dragon2 {
	return &Dragon2{}
}

type Shape interface {
	Render() string
}

type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle of radius %f", c.Radius)
}

func (c *Circle) Resize(factor float32) {
	c.Radius *= factor
}

type Square struct {
	Side float32
}

func (s *Square) Render() string {
	return fmt.Sprintf("Square with side %f", s.Side)
}

type ColoredShape struct {
	Shape Shape
	Color string
}

func (c *ColoredShape) Render() string {
	return c.Shape.Render() + " has the color " + c.Color
}

type TransparentShape struct {
	Shape        Shape
	Transparency float32
}

func (t *TransparentShape) Render() string {
	return fmt.Sprintf("%s has %f%% transparency", t.Shape.Render(), t.Transparency*100.0)
}

func main() {
	d := Dragon{}
	d.Bird.Age = 11
	d.Fly()
	d.Lizard.Age = 9
	d.Crawl()
	d.Fly()

	d2 := NewDragon2()
	d2.SetAge(11)
	d2.Fly()
	d2.SetAge(9)
	d2.Crawl()
	d2.Fly()

	c := Circle{2}
	fmt.Println(c.Render())
	c.Resize(2)
	fmt.Println(c.Render())

	redSquare := ColoredShape{&Square{5}, "Red"}
	fmt.Println(redSquare.Render())

	redHalfTransparentCircle := TransparentShape{&ColoredShape{&Circle{5}, "Red"}, 0.5}
	fmt.Println(redHalfTransparentCircle.Render())
}
