package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Running %s go on %s\n", os.Args[0], os.Getenv("GOFILE"))
	fmt.Printf("  os.Args = %#v\n", os.Args)

	if len(os.Args) < 2 {
		fmt.Println("No argument provided. Exiting")
		return
	}

	if os.Args[1] == "manufacturers" {
		fmt.Println("Generating internal/bt/manufacturers.go ...")
		generateManufacturers()
	} else if os.Args[1] == "services" {
		fmt.Println("Generating internal/bt/services.go ...")
		generateServices()
	}
}
