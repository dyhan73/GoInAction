package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type ID int
type Manager struct {
	ID
	Employee // embedding - only type without variable name
	Reports  []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	return m.Reports
}

func embeddingTest() {
	m := Manager{
		ID: 999,
		Employee: Employee{
			Name: "Bob Bobson",
			ID:   "12345",
		},
		Reports: []Employee{},
	}
	fmt.Println(m.Name, m.ID, m.Employee.ID)
	fmt.Println(m.Description(), m.FindNewEmployees())

	//var eFail Employee = m // Manager can't be converted to Employee
	var eOK Employee = m.Employee
	fmt.Println(eOK)
}

func main3() {
	embeddingTest()
}
