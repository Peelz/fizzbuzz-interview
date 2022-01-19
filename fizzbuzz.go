package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Err() != nil {
		// Handle error.
	}
	inp := scanner.Text()
	// Do something here
	fmt.Println(inp)
}
