package main

import(
	"fmt"
)

// constant variable declaration
const constant int = 90

// function that takes 2 arguments and has return type of float64
func add(x float64, y float64) float64{

	return x+y
}

// This is short cut way of the add function
//func add(x, y float64) float64{
//
//	return x+y
//}

/*
This function will return multiple values
 */
func multipleReturns(x, y string) (string, string){

	return x, y
}

func main(){

	var num1 float64 = 11.56
	var num2 float64 = 4.27

	// short cut way of variables decparation
	//num1, num2 := 11.56, 4.27

	fmt.Printf("Constant value is : %v\n", constant)

	fmt.Printf("Summation of %v and %v is : %v\n\n", num1, num2, add(num1, num2))

	str1, str2 := "Yaaoo", "Rev"
	fmt.Println(multipleReturns(str1, str2))

	var rand1 int = 10
	// type casting
	var rand2 float64 = float64(rand1)
	//x := rand1	// x will be type int

	fmt.Printf("\nInt value : %v\n", rand1)
	fmt.Printf("Floating value : %v\n", rand2)
}
