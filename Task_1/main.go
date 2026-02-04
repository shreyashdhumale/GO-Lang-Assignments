package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Name input
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// Age input
	fmt.Print("Enter your age: ")
	ageInput, _ := reader.ReadString('\n')
	ageInput = strings.TrimSpace(ageInput)

	age, err := strconv.Atoi(ageInput)
	if err != nil {
		fmt.Println("Age must be a valid number")
		return
	}

	// Negative age check
	if age < 0 {
		fmt.Println("Age cannot be negative")
		return
	}

	person := Person{
		Name: name,
		Age:  age,
	}

	person.Introduce()
	fmt.Println("Eligible to vote:", person.IsEligibleToVote())
}
