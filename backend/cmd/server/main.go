package main

import (
	"fmt"
	"os"
)

func main() {
	if err := Execute(); err != nil {
		fmt.Printf("execution failed: %s\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
