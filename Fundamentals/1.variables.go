package main

import "fmt"

func main() {
	fmt.Println("Hello Dinuka")

	var game = "last of us" //declaring and assigning variable
	//game := "last of us" we can declare and assign value like this without using var keyword
	fmt.Println(game)

	game = "last of us 2" //reassigning
	fmt.Println(game)

	//game = 12 (not working becase go is typed language)

	var game2 string //we can define variable like this and assign value later and when we assign we can assign different data type values
	game2 = "apex ledgends"
	fmt.Println(game2)

	//declaring multiple variables
	//Compound creation
	var var1, var2, var3 = 28, "Dota2", "Golang"
	//var1, var2, var3 := 28, "Dota2", "Golang" withou var keyword
	//var1 := "league of ledgends" we can't reassign like this but we can reassign along with new variable ex- var4, var1 := "new var", "CoD"
	fmt.Println(var1, var2, var3)

	//Block creation
	var (
		var4 = 26
		var5 = "DSW"
	)
	fmt.Println(var4, var5)

	//CONSTANTS
	const myName = " Dinuka "

}
