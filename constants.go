package main

import ( 
	"fmt"
	"math"
)

//Declaring a constant
//Constant can appear anywhere where Variable also can appear
const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000
	
	//A constant do an arithmathic expression with random precision
	const d = 3e20 / n
	fmt.Println(d)

	//Numeric constant doesn't have a type until given one
	//For example by converting a number to a integer 64 bit type
	fmt.Println(int64(d))

	//Or by using context like in a function
	//Here the math.Sin function expect a float64 input
	//Thus, making the constant n a float64 type
	fmt.Println(math.Sin(n))
}
