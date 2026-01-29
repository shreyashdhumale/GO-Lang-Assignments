package main

import "fmt"

// Person Struct
type Person struct {
	Name string
	Age  int
}

// Introduction
func (p Person) Introduce() {
	fmt.Printf("Hi, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

// Age Update
func (p *Person) UpdateAge(newAge int) {
	p.Age = newAge
}

// Vote check
func (p Person) IsEligibleToVote() bool {
	return p.Age >= 18
}

func main() {
	person := Person{
		Name: "Shreyash",
		Age:  21,
	}
	person.Introduce() // Calling Intro

	person.UpdateAge(22) // Age update

	person.Introduce() // Checking the Updated age

	fmt.Println("Eligible to vote", person.IsEligibleToVote())
}