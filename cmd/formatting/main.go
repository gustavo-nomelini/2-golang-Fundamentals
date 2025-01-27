package main

import "fmt"

type ID int

var (
	car        string  = "onix"
	year       int     = 2020
	carMileage float64 = 37350.73
	carID      ID      = 1
)

func main() {
	fmt.Printf("The type of the variable CAR is %T\n", car)
	fmt.Printf("The name of the car is: %v\n", car)
	fmt.Printf("The car mileage is %v\n", carMileage)
	fmt.Printf("The car mileage is %f\n", carMileage)
	// Use fmt.Printf when you need control over the formatting,
	// and fmt.Println for simpler cases where formatting isn't needed.
	fmt.Println("Car: ", car, year)
	//---------------------------------
	fmt.Printf("The model of the car is %v and the year of the model is %v\n", car, year)
}
