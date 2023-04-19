package main

import "fmt"

func main() {
	//Arithmatic operators
	// + - / * %
	var n1 = 3
	var n2 = 4

	var sum = n1 + n2
	var remainder = n1 % 2
	fmt.Println(sum, remainder)

	//Relational operators > < >= <= == !=
	var results = n1 == n2
	fmt.Println(results)

	//Logical operators &&(AND), ||(OR), !(Not)
	const city = "Matara"
	var age = 26

	var n3 = !(city == "Matara") || (age >= 25 && age < 30)
	fmt.Println(n3)
}
