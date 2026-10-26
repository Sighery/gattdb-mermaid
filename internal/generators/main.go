package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No argument provided. Exiting...")
		return
	}

	switch os.Args[1] {
	case "manufacturers":
		fmt.Println("Generating internal/bt/manufacturers.go ...")
		generateManufacturers()
	case "services":
		fmt.Println("Generating internal/bt/services.go ...")
		generateServices()
	}
}
