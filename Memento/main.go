package main

import "fmt"

type BankAccount struct {
	balance int
}

func (b *BankAccount) Deposit(amount int) *Memento {
	b.balance += amount
	return &Memento{b.balance}
}

func (b *BankAccount) Restore(m *Memento) {
	b.balance = m.balance
}

type Memento struct {
	balance int
}

func (b *BankAccount) String() string {
	return fmt.Sprintf("Balance: %d", b.balance)
}

type BankAccountM struct {
	balance int
	changes []*Memento
	current int
}

func (b *BankAccountM) Deposit(amount int) *Memento {
	b.balance += amount
	m := &Memento{b.balance}
	b.changes = append(b.changes, m)
	b.current++
	return m
}

func (b *BankAccountM) Undo() *Memento {
	if b.current > 0 {
		b.current--
		m := b.changes[b.current]
		b.balance = m.balance
		return m
	}
	return nil
}

func main() {
	ba := BankAccount{100}
	m1 := ba.Deposit(50)
	_ = ba.Deposit(25)
	fmt.Println(ba.String())
	ba.Restore(m1)
	fmt.Println(ba.String())

}
