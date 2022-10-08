package main

type Employee1 struct {
	Name, Position string
	AnnualIncome   int
}

const (
	Developer = iota
	Manager
)

type EmployeeFactory1 struct {
	Position     string
	AnnualIncome int
}

func NewEmployee(role int) *Employee1 {
	switch role {
	case Developer:
		return &Employee1{"", "Developer", 60000}
	case Manager:
		return &Employee1{"", "Manager", 80000}
	default:
		panic("unsupported role")
	}
}
