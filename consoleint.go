package  main

import (
	  "fmt"
	  "os"
	  "bufio"
	  "strconv"
)

func main()  {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter your year :")
	scanner.Scan()
	input, _:= strconv.ParseInt(scanner.Text(), 10,64)
	fmt.Printf("You are %d  years old at  the end of 2020", 2020-input)
	
}