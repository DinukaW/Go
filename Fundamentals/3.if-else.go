package main

import "fmt"

func main() {
	age := 30

	if age >= 20 && age < 25 {
		fmt.Println("You're not Invited! too young")
	} else if age >= 25 && age <= 30 {
		fmt.Println("You're Invited!")
	} else {
		fmt.Println("You're not Invited!")
	}

}
