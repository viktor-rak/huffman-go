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
	frequency := make(map[rune]int)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter a string you would like to encode: ")
	scanner.Scan()
	input := scanner.Text()
	for _, freq := range input {
		frequency[freq]++
	}
	fmt.Println("Input: ", input)
	fmt.Println(frequency)
}
