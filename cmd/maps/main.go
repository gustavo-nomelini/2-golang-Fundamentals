// 2025-01-28 14:52
// Gustavo Lopes Nomelini
// MAPS in GOLANG works like HASHTABLES in another languages
// they are used to store data values in key:value pairs

package main

import "fmt"

func main() {
	// Create a map with string keys and int values
	m := map[string]int{
		"apple":  5,
		"banana": 3,
	}

	// Add a new key-value pair
	m["orange"] = 2

	// Access a value
	fmt.Println("Apple count:", m["apple"]) // Output: Apple count: 5

	// Delete a key
	delete(m, "banana")

	fmt.Println(m) // Output: map[apple:5 orange:2]

	// initializing an empty map (nil values)
	test := make(map[string]string)
	test["Gustavo"] = "programador"

	fmt.Println(test)
}
