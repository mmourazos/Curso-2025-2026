package main

import "fmt"

func main() {
	const Pi = 3.14
	s := []int{1, 2, 3, 4, 5}

	s = append(s, 1, 2, 3)

	a := [...]int{1, 2, 3, 4, 5}

	fmt.Print("Hello world!", Pi, a, s)
}
