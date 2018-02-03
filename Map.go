package main

import "fmt"

func main(){

	// it's a map which's key is a string type and value is float type
	//var grades map[string] float32

	// to have values we need to use make()
	grades := make(map[string]float32)	// so we can now add values to it and also can get values from it

	grades["John"] = 91
	grades["Murdock"] = 56
	grades["Frank"] = 74

	fmt.Println(grades)

	JohnsGrade := grades["John"]
	MurdocksGrade := grades["Murdock"]
	FranksGrade := grades["Frank"]

	fmt.Printf("\n%v\n%v\n%v\n", JohnsGrade, MurdocksGrade, FranksGrade)

	delete(grades, "John")

	fmt.Println()
	fmt.Println(grades)
	fmt.Println()

	for k, v := range grades{

		fmt.Println(k, "", v)
	}
}
