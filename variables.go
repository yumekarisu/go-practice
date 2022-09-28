package main

import "fmt"

func main() {
	//Declaring a Variable
	var a = "initial"
	fmt.Println(a)

	//You can declare multiple varible at once
	//In Go, the variable name go first, then the type
	var b, c int = 1, 2
	fmt.Println(b, c)

	////Go will automatically infer the type of initialized variable
	var d = true
	fmt.Println(d)

	//Declared variable without an initial value is gonna be zero in value
	//For example, int type variable without initial value is gonna have
	//a value of 0
	var e int
	fmt.Println(e)

	//use the := to quickly create a varible
	//Go will automatically infer the type based on the entered value
	f := "apple"
	fmt.Println(f)
}
