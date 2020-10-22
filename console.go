package main

import (
	"os"
	"fmt"
	"bufio"
)

func main ()  {
		scanner:= bufio.NewScanner(os.Stdin)
		fmt.Printf("Enter the name:")
		scanner.Scan()
		input :=scanner.Text()
		fmt.Printf("Your name: %d", input)
	
}