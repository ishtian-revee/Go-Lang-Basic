package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){

	fmt.Println("Before seeding random number", rand.Intn(100))

	// random number generator is deterministic in go lang
	// that is why it will produce the same number sequence of each execution of the program
	fmt.Println("\nWith loop (20 random numbers): ")
	for i:=0; i<20; i++ {

		fmt.Printf("%v. Value = %v\n", i+1, rand.Intn(20))
	}

	fmt.Printf("\nFloating random number (float32): %v\n", rand.Float32())
	fmt.Printf("Floating random number (float64): %v\n\n", rand.Float64())

	// seeding
	// it will generate random number sequence each time of execution of the program
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	for i:=0; i<20; i++ {

		fmt.Printf("%v. Value = %v\n", i+1, random.Intn(20))
	}
}
