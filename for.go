package main

import "fmt"

func main() {

	//For is the only looping keyword in go
	//(there's no dedicated keyword for while)
	//However, you can sorta do while loop 
	//by using the most basic for loop in go
	//The one with only a single condition
	i := 1
	for i < 3 {
		fmt.Println(i)
		i = i + 1
	}

	//Standard For loop in Go
	//Remember to use semicolon (;) to separate
	//the initial, condition, and after statement
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	//You can also declare an infinite loop
	//that can only be escaped by break or return
	for {
		fmt.Println("loop")
		break
	}

	//You can also continue to the next iteration of the loop
	//essentially skiping the current loop if the condition is matched
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

