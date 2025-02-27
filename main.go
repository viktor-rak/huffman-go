package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	frequency int
	left      *Node
	right     *Node
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter a string you would like to encode: ")
	scanner.Scan()
	input := scanner.Text()
	fmt.Println("Input: ", input)
}
