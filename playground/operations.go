package main

import "fmt"

func main() {
	// var (
	// 	name1 = "test"
	// 	name2 = "test"
	// )
	// var result bool = name1 != name2
	// fmt.Println(result)

	var (
		True = true;
		False = false;
	)
	fmt.Println(True && False) // false
	fmt.Println(True || False) // true
	fmt.Println(!False) // true, it's a unary operator
}