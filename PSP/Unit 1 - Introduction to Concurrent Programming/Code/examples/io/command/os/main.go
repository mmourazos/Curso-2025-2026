package main

import (
	"fmt"
	"os"
	"time"
)

func age(birthdate time.Time) int {
	today := time.Now()
	age := today.Year() - birthdate.Year()

	if birthdate.Month() > today.Month() {
		age--
	}

	return age
}

func main() {
	// os.Args is a slice of strings that holds the command-line arguments passed to the program.

	// We can check the number of arguments. Let's say we want to get something like:
	// os.exe -n <name> -b <birth date>
	// We are expecting 5 arguments:
	// * First argument is the name of the command ".../os.x".
	// * Second arg.: -n.
	// * Third arg. : the name.
	// * Forth arg. : -b
	// * Fifth arg. : the date. 15/05/1997
	if len(os.Args) != 5 {
		fmt.Print("Usage: os.exe -n <name> -b <birth date>")
	} else {
		name := os.Args[2]
		birthStr := os.Args[4]
		birthDate, _ := time.Parse("02/01/2006", birthStr)

		myAge := age(birthDate)
		fmt.Printf("Your name is %s and you age is %d.\n", name, myAge)
	}
}
