package main

import "fmt"

func compare(value int) string {
	//do not change this variable resultMessage, secretValue
	resultMessge := ""
	secretValue := 88

	//Insert your code from here
	if value == secretValue {
		fmt.Println("Well done, Correct Guess")
	}
	if value < secretValue {
		fmt.Println("Too Low, try again")
	}
	if value > secretValue {
		fmt.Println("Too high, try again")
	}
	//do not remove this line
	return resultMessge
}

func main() {
	var guess int
	fmt.Println("Enter an integer value: ")
	fmt.Scanln(&guess)
	compare(guess)
}
