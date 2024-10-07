package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Poke-CLI > ")
	scanner.Scan()
	data := scanner.Text()
	fmt.Println(data)
}
