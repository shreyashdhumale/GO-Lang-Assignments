package main

import "fmt"

// Person struct
type Person struct {
	Name string
	Age  int
}

// Introduction method
func (p Person) Introduce() {
	fmt.Printf("Hi, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

// Vote check
func (p Person) IsEligibleToVote() bool {
	return p.Age >= 18
}
