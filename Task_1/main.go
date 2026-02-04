package main

import "fmt"

func main() {
	person := Person{
		Name: "Shreyash",
		Age:  21,
	}

	person.Introduce()

	// Direct assignment (no method needed)
	person.Age = 22

	person.Introduce()

	fmt.Println("Eligible to vote:", person.IsEligibleToVote())
}
