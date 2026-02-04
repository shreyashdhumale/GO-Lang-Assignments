package main

import "fmt"

func main() {

	ramesh := &Employee{"Ramesh", 25, 50000}
	suresh := &Employee{"Suresh", 28, 60000}

	itDepartment := Department{Name: "IT"}

	itDepartment.AddEmployee(ramesh)
	itDepartment.AddEmployee(suresh)

	fmt.Println("Current Average Salary:", itDepartment.AverageSalary())

	// Direct salary update through behavior
	ramesh.GiveRaise(5000)

	fmt.Println("Average Salary after raise:", itDepartment.AverageSalary())

	// Direct removal 
	for i, e := range itDepartment.Employees {
		if e.Name == "Ramesh" {
			itDepartment.Employees =
				append(itDepartment.Employees[:i], itDepartment.Employees[i+1:]...)
			fmt.Println("Ramesh has left the IT department")
			break
		}
	}

	fmt.Println("Final Average Salary:", itDepartment.AverageSalary())
}
