package main

import "fmt"

type Employee struct {
	Name   string
	Age    int
	Salary float64
}

type Department struct {
	Name      string
	Employees []*Employee
}

func (e *Employee) GiveRaise(amount float32) {
 	fmt.Printf("Giving ₹%.0f raise to %s\n", amount, e.Name)
	e.Salary += float64(amount)
}

// Add Employee 
func (d *Department) AddEmployee(e *Employee) {
	fmt.Printf("%s joined the %s department\n", e.Name, d.Name)
	d.Employees = append(d.Employees, e)
}

// Remove Employee
func (d *Department) RemoveEmployee(name string) {
	for i, e := range d.Employees {
		if e.Name == name {
			fmt.Printf("%s has left the %s department\n", e.Name, d.Name)
			d.Employees = append(d.Employees[:i], d.Employees[i+1:]...)
			return
		}
	}
}

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

func main() {

	ramesh := &Employee{"Ramesh", 25, 50000}
	suresh := &Employee{"Suresh", 28, 60000}

	itDepartment := Department{Name: "IT"}

	itDepartment.AddEmployee(ramesh)
	itDepartment.AddEmployee(suresh)

	fmt.Println("Current Average Salary:", itDepartment.AverageSalary())

	ramesh.GiveRaise(5000)

	fmt.Println("Average Salary after raise:", itDepartment.AverageSalary())

	itDepartment.RemoveEmployee("Ramesh")

	fmt.Println("Final Average Salary:", itDepartment.AverageSalary())
}
