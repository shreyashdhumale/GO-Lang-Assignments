package main

import "fmt"

// Employee struct
type Employee struct {
	Name   string
	Age    int
	Salary float64
}

// Department struct
type Department struct {
	Name      string
	Employees []*Employee
}

// GiveRaise 
func (e *Employee) GiveRaise(amount float64) {
	fmt.Printf("Giving ₹%.0f raise to %s\n", amount, e.Name)
	e.Salary += amount
}

// AddEmployee 
func (d *Department) AddEmployee(e *Employee) {
	fmt.Printf("%s joined the %s department\n", e.Name, d.Name)
	d.Employees = append(d.Employees, e)
}

// AverageSalary 
func (d Department) AverageSalary() float64 {
	if len(d.Employees) == 0 {
		return 0
	}

	total := 0.0
	for _, e := range d.Employees {
		total += e.Salary
	}
	return total / float64(len(d.Employees))
}
