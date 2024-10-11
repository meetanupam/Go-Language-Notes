package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	welcome:= "welcome to user input"
	fmt.Println(welcome)

	reader:= bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for out pizza:")

	// hello := "to here"
	// fmt.Println(hello)

	// comma om and err (error) ok syntax

	input, _ := reader.ReadString('\n')
	// input and hold an error
	fmt.Println("Thanks for rating ", input)
	fmt.Printf("Thanks for rating %T \n", input)

	// when the majority of focus in on error
	// _, err := reader.ReadString('\n')
	// fmt.Println("Thanks for rating ", err)

	

}