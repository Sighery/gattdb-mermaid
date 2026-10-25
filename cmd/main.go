package main

import (
	"os"

	"github.com/Sighery/gattdb-mermaid/internal/bt"
)

func main() {
	if len(os.Args) >= 2 && os.Args[1] == "scan" {
		bt.Scan()
	}
}
