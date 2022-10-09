package main

import "fmt"

type Creature struct {
	Name            string
	Attack, Defense int
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attack, c.Defense)
}

func NewCreature(name string, attack, defense int) *Creature {
	return &Creature{name, attack, defense}
}

type Modifier interface {
	Add(m Modifier)
	Handle()
}

type CreatureModifier struct {
	creature *Creature
	next     Modifier
}

func (c *CreatureModifier) Add(m Modifier) {
	if c.next != nil {
		c.next.Add(m)
	} else {
		c.next = m
	}
}

func (c *CreatureModifier) Handle() {
	if c.next != nil {
		c.next.Handle()
	}
}

func NewCreatureModifier(c *Creature) *CreatureModifier {
	return &CreatureModifier{creature: c}
}

type DoubleAttackModifier struct {
	CreatureModifier
}

func NewDoubleAttackModifier(c *Creature) *DoubleAttackModifier {
	return &DoubleAttackModifier{CreatureModifier{creature: c}}
}

func (d *DoubleAttackModifier) Handle() {
	fmt.Println("Doubling", d.creature.Name, "'s attack")
	d.creature.Attack *= 2
	d.CreatureModifier.Handle()
}

func main() {
	goblin := NewCreature("Goblin", 1, 1)

	fmt.Println(goblin.String())

	root := NewCreatureModifier(goblin)

	root.Add(NewDoubleAttackModifier(goblin))
	root.Add(NewDoubleAttackModifier(goblin))

	root.Handle()

	fmt.Println(goblin.String())
}
