package main

const pi float64 = 3.14

// declaring global variables
var (
	name   string  = "Gustavo"
	age    int     = 26
	weight float64 = 65.3
)

// :=  only works on the function scope
// test int := 3

// using TYPES in golang
type ID int

func main() {
	println(pi)
	println(name)
	println(age)
	println(weight)

	// have to declare inside main if using :=
	// short declaration, the type is alwas inferred from the value
	test1 := 3
	println(test1)

	// when using := we can declare a new value for the variable
	// ONLY IF we declare a new variable
	test1, test2 := 10, 15
	println(test1, test2)

	var age int = 27
	println(age)

	// initialize the variable using :=
	// change it using =
	test3 := 30
	test3 = 37
	println(test3)
}
