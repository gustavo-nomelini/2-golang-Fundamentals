package main

import "fmt"

const numberBands = 4

func main() {
	var games [3]string
	games[0] = "Ragnarok"
	games[1] = "World of Warcraft"
	games[2] = "Baldurs Gate 3"

	// fmt.Printf(games[0])

	for index, name := range games {
		fmt.Printf("The game in the array index %v is %v\n", index, name)
	}

	println("----------------------------------------")

	bands := [numberBands]string{
		"System Of A Down",
		"Linkin Park",
		"Limp Bizkit",
		"Iron Maiden",
	}

	for i := 0; i < numberBands; i++ {
		fmt.Printf("The band in the array index %v is %v\n", i, bands[i])
	}

	println("----------------------------------------")

	fruits := [...]string{"Apple", "Banana", "Peach"}
	// cant carry the index
	for _, v := range fruits {
		fmt.Println(v)
	}
}
