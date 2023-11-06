package main

import (
	"fmt"
	"os"

	"github.com/fourhundredfour/brcmp/pkg/comparison"
)

var (
	firstBinaryPath  string
	secondBinaryPath string
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: brcmp <binary1> <binary2>")
		os.Exit(1)
	}
	firstBinaryPath = os.Args[1]
	secondBinaryPath = os.Args[2]
	firstBinary, err := os.ReadFile(firstBinaryPath)
	secondBinary, err := os.ReadFile(secondBinaryPath)

	if err != nil {
		panic(err)
	}
	fmt.Print(comparison.CompareBinaries(firstBinary, secondBinary))
	os.Exit(0)
}
