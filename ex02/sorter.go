package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/tungtanium/go-course23/ex02/sortitout"
)

func main() {
	// Flags
	intFlag := flag.Bool("int", false, "Sort integer data increasingly. Data to sort must be in numeric. E.g: 1 2 3 4")
	floatFlag := flag.Bool("float", false, "Sort float data increasingly. Data to sort must be in numeric. E.g: 1.0 2.1 3.2 4.3")
	stringFlag := flag.Bool("string", false, "Sort string data alphabetically. Data to sort must be in non-numeric. E.g: Mot hai ba bon")

	flag.Parse()

	inputData := flag.Args()
	inputFlag := flag.NFlag()

	if inputFlag > 1 {
		log.Panic("Một flag thôi cha!!! Tui mới học")
		panic(inputFlag)
	}

	if *intFlag {
		fmt.Println("Result: ", sortitout.SortInt(inputData))
	} else if *floatFlag {
		fmt.Println("Result: ", sortitout.SortFloat(inputData))
	} else if *stringFlag {
		fmt.Println("Result: ", sortitout.SortString(inputData))
	}
}
