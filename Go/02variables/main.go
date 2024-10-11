package main

import "fmt"

const loginToken string = "abcdef" // public

func main(){
	// fmt.Println("variables");
	var username string = "variables" // string
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)
	
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	var smallVal uint8 = 255 // uint8 can only stores the value upto 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type : %T \n", smallVal)

	var smallFloat float32 = 255.4554442421
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)

	var longFloat float64 = 255.4554442421
	fmt.Println(longFloat)
	fmt.Printf("Variable is of type : %T \n", longFloat)

	// default value and some alises
	var anotherVariable int // if u declare an integer it is going to be '0'
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type : %T \n", anotherVariable)

	// implicit type and way of decalaring the variables
	var website = "anupamshakya.com"
	fmt.Println(website)

	// Another way of declaring variables inside a method or function.
	numberTotal := 3000
	fmt.Println(numberTotal)

	fmt.Println(loginToken)
	fmt.Printf("Variable is of type : %T \n", loginToken)

	// var age int = 23 // number or integer
	// var condition bool = true
	// var fll float64 = 100.100
	
	// fmt.Println(age) // printing age
	// fmt.Println(condition) // condition
	// fmt.Print(fll) // float
}