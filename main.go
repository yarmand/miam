package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, world:" + os.Getenv("LAPIN"))
}
