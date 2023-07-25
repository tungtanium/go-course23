package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/tungtanium/go-course23/ex02/sortitout"
)

func main() {
	// Flags
	intFlag := flag.Bool("int", false, "Sort integer data increasingly")
	floatFlag := flag.Bool("float", false, "Sort float data increasingly")
	stringFlag := flag.Bool("string", false, "Sort string data alphabetically")

	// flag.Func example
	// var fooStr []string
	// flag.Func("foo", "Test custom flag named Foo", func(flagValue string) error {
	// 	fooStr = strings.Fields(flagValue)
	// 	return nil
	// })
	// intData := flag.Bool("int", false, "Sort int data ascendingly")
	// floatData := flag.Bool("float", false, "Sort float data ascendingly")

	flag.Parse()

	// foo := flag.Set()
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

	// fmt.Println("Tails: ", inputData)
	// fmt.Println("Flag -int: ", *intFlag)
	// fmt.Println("Flag -float: ", *floatFlag)
	// fmt.Println("Flag -string: ", *stringFlag)
	// fmt.Println("Foo flag: ", fooStr)
	// fmt.Println("Flag float: ", *floatData)
	// fmt.Println("Num of Flags: ", inputFlag)
	// fmt.Println("Num of Args: ", numArgs)

	// Test import
	// sortitout.Testing()
}
