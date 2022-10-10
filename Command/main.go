package main

import "fmt"

var overdraftLimit = -500

type BankAccount struct {
	balance int
}

func (b *BankAccount) Deposit(amount int) bool {
	b.balance += amount
	fmt.Println("Deposited", amount, "balance is now", b.balance)
	return true
}

func (b *BankAccount) Withdraw(amount int) bool {
	if b.balance-amount > overdraftLimit {
		b.balance -= amount
		fmt.Println("Withdrew", amount, "balance is now", b.balance)
		return true
	}
	fmt.Println("Cannot withdraw", amount, "balance is only", b.balance)
	return false
}

func (b *BankAccount) Balance() int {
	return b.balance
}

func (b *BankAccount) OverdraftLimit() int {
	return overdraftLimit
}

func (b *BankAccount) SetOverdraftLimit(limit int) {
	overdraftLimit = limit
}

type Command interface {
	Call()
	Undo()
	Succeeded() bool
	SetSucceeded(v bool)
}

type BankAccountCommand struct {
	account *BankAccount
	action  Action
	amount  int
	success bool
}

type Action int

const (
	Deposit Action = iota
	Withdraw
)

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{account, action, amount, false}
}

func (c *BankAccountCommand) Call() {
	switch c.action {
	case Deposit:
		c.success = c.account.Deposit(c.amount)
	case Withdraw:
		c.success = c.account.Withdraw(c.amount)
	}
}

type CompositeBankAccountCommand struct {
	commands []Command
}

func NewCompositeBankAccountCommand(commands []Command) *CompositeBankAccountCommand {
	return &CompositeBankAccountCommand{commands}
}

func (c *CompositeBankAccountCommand) Call() {
	for _, cmd := range c.commands {
		cmd.Call()
	}
}

func (c *CompositeBankAccountCommand) Undo() {
	for i := len(c.commands) - 1; i >= 0; i-- {
		c.commands[i].Undo()
	}
}

func (c *CompositeBankAccountCommand) Succeeded() bool {
	for _, cmd := range c.commands {
		if !cmd.Succeeded() {
			return false
		}
	}
	return true
}

func (c *CompositeBankAccountCommand) SetSucceeded(v bool) {
	for _, cmd := range c.commands {
		cmd.SetSucceeded(v)
	}
}

type MoneyTransferCommand struct {
	CompositeBankAccountCommand
	from, to *BankAccount
	amount   int
}

func NewMoneyTransferCommand(from, to *BankAccount, amount int) *MoneyTransferCommand {
	c := NewCompositeBankAccountCommand([]Command{
		NewBankAccountCommand(from, Withdraw, amount),
		NewBankAccountCommand(to, Deposit, amount),
	})
	return &MoneyTransferCommand{*c, from, to, amount}
}

func (c *MoneyTransferCommand) Call() {
	ok := c.from.Withdraw(c.amount)
	if !ok {
		return
	}
	c.to.Deposit(c.amount)
	c.SetSucceeded(true)
}

func (c *MoneyTransferCommand) Undo() {
	if c.Succeeded() {
		c.CompositeBankAccountCommand.Undo()
	}
}

func (c *MoneyTransferCommand) Succeeded() bool {
	return c.CompositeBankAccountCommand.Succeeded()
}

func (c *MoneyTransferCommand) SetSucceeded(v bool) {
	c.CompositeBankAccountCommand.SetSucceeded(v)
}

func (c *BankAccountCommand) Undo() {
	if !c.success {
		fmt.Println("Cannot undo")
		return
	}

	switch c.action {
	case Deposit:
		c.account.Withdraw(c.amount)
	case Withdraw:
		c.account.Deposit(c.amount)
	}
}

func (c *BankAccountCommand) Succeeded() bool {
	return c.success
}

func (c *BankAccountCommand) SetSucceeded(v bool) {
	c.success = v
}

func main() {
	ba := BankAccount{}
	cmd := NewBankAccountCommand(&ba, Deposit, 100)
	cmd.Call()
	fmt.Println(ba)

	cmd = NewBankAccountCommand(&ba, Withdraw, 50)
	cmd.Call()
	fmt.Println(ba)

	cmd.Undo()
	fmt.Println(ba)

	cmd = NewBankAccountCommand(&ba, Withdraw, 5000)
	cmd.Call()
	fmt.Println(ba)

	cmd.Undo()

	mtc := NewMoneyTransferCommand(&ba, &BankAccount{}, 100)
	mtc.Call()
	fmt.Println(ba)

	mtc.Undo()
	fmt.Println(ba)

}
