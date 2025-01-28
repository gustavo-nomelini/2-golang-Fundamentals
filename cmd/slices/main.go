package main

import "fmt"

func main() {
	names := [4]string{
		"Gustavo",
		"Jessica",
		"Rafael",
		"Bruna",
	}
	fmt.Println(names)
	fmt.Println("--------------------------------")

	slice1 := names[1:3]
	fmt.Println(slice1)

	slice2 := names[:2]
	fmt.Println(slice2)

	slice3 := names[2:]
	fmt.Println(slice3)

	// create a slice from an array
	// referencing the storage
	slice4 := names[:]
	fmt.Println(slice4)

	fmt.Println("--------------------------------")

	fmt.Printf("len=%d cap=%d %v\n", len(slice1), cap(slice1), slice1)
	fmt.Printf("len=%d cap=%d %v\n", len(slice2), cap(slice2), slice2)
	fmt.Printf("len=%d cap=%d %v\n", len(slice3), cap(slice3), slice3)
	fmt.Printf("len=%d cap=%d %v\n", len(slice4), cap(slice4), slice4)
}
