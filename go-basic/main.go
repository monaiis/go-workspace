package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Variable declaration

	var x int = 10
	y := 20

	fmt.Println(x, y)

	// Array & Slices(list)

	var z [5]int = [5]int{1, 2, 3, 4, 5}
	w := [5]int{1, 2, 3, 4, 5}
	q := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(z[3:], w, q)

	slice := []int{1, 2}
	slice = append(slice, 3)
	// If we use undeclare variable, we must use :=
	new_slice := append(slice, 4)
	l := len(new_slice)

	fmt.Printf("%#v\n%#v, len = %v\n\n", slice, new_slice, l)

	// 'len()' can use in case to find length of array/slice but not work on string in some case (languanges)
	// If we want to count length of string, we have to use utf8 pakage

	name := "John"
	name_length := utf8.RuneCountInString(name)

	fmt.Printf("%v, len = %v\n\n", name, name_length)

	// Map (dict) -> map[key]value

	countries := map[string]string{} // var countries map[string]string
	countries["th"] = "Thailand"
	countries["en"] = "English"

	fmt.Printf("%v", countries)

	country, ok := countries["jp"] // Map return two values which is value and boolean

	if ok {
		println(country)
	} else {
		println("no value\n")
	}

	// Loop

	values := []int{10, 20, 30, 40}

	for i := 0; i < len(values); i++ {
		fmt.Printf("%v ", values[i])
	}
	println()

	// Go doesn't have 'while' to conserve the reserved word
	// 'while' have to write in this form instead

	j := len(values) - 1
	for j >= 0 {
		fmt.Printf("%v ", values[j])
		j--
	}
	println()

	// for each

	for index, value := range values { // If we want to ignore index or value we can use _ instead
		println(index, value)
	}

	// Function

	result := add(10, 20)
	fmt.Printf("\n%v\n", result)
	hello("Doe")
	// Anonymous function
	fn := func(a, b int) int {
		return a + b
	}
	sum := fn(10, 30)
	println(sum)

	cal(add)

}

func cal(f func(int, int) int) {
	sum := f(20, 30)
	println(sum)
}

func add(a, b int) int { // We can return many values using tuple -> (int, string, bool)
	return a + b
}

func hello(name string) {
	println("Hello " + name)
}
