package main

import "fmt"

const usicteenbitmax float64 = 65535
const kph_multiple float64 = 1.60934

// structure type
type car struct{
	gas_padal uint16
	brake_padal uint16
	steering_wheel int16
	top_speed float64
}

/*
value receiver method
this can't change the values of struct properties
it generally copies the value of whatever passes as the arguments
  */
func (c car) kph() float64{

	return float64(c.gas_padal) * (c.top_speed/usicteenbitmax)
}


// value receiver method
func (c car) mph() float64{

	return float64(c.gas_padal) * (c.top_speed/usicteenbitmax/kph_multiple)
}

/*
pointer receiver method
this can change the values of struct properties
is generally modifies the value of whatever passes as the arguments
  */
func (c *car) new_top_speed(newSpeed float64){

	c.top_speed = newSpeed
}

func newer_top_speed(c car, speed float64) car{

	c.top_speed = speed
	return c
}

func main(){

	a_car := car{
		gas_padal: 22450,
		brake_padal: 0,
		steering_wheel: 12000,
		top_speed: 250.0,
	}

	//a_car := car{22450, 0, 12000, 250.0}

	// value receiver method calls
	fmt.Println("Gas padal: ", a_car.gas_padal)
	fmt.Println("Kilometers: ", a_car.kph())
	fmt.Println("Miles: ", a_car.mph())

	// pointer receiver method call
	a_car.new_top_speed(380.0)
	fmt.Println("\nKilometers: ", a_car.kph())
	fmt.Println("Miles: ", a_car.mph())

	// general function call
	a_car = newer_top_speed(a_car, 490.0)
	fmt.Println("\nKilometers: ", a_car.kph())
	fmt.Println("Miles: ", a_car.mph())
}
