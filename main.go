package main

import (
	"fmt"
	"os"
)

func initArbiter() *Arbiter {
	a := NewArbiter()
	a.InitManufacturers()
	return a
}

func printManufacturers(man []string) {
	fmt.Println("Supported manufacturers:")
	for _, m := range man {
		fmt.Printf("  - %s\n", m)
	}
}

func printParts(man string, parts []string) {
	fmt.Printf("Parts list for %s\n", man)
	for _, p := range parts {
		fmt.Printf("  - %s\n", p)
	}
}

func main() {
	a := initArbiter()

	mfs := a.ListManufacturers()

	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	if os.Args[1] == "list" {
		printManufacturers(mfs)
		os.Exit(0)
	}

	mfName := os.Args[1]

	m, err := a.GetManufacturer(mfName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	if len(os.Args) < 3 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	if os.Args[2] == "list" {
		printParts(mfName, m.ListParts())
	}
}
