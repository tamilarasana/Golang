package main

import "fmt"

func main()  {
	ans := 25
	switch  {
	case ans < 20:
		fmt.Println("Less Then 20")
	case ans > 20:
		fmt.Println("Greater Then 20")		
	default:
		fmt.Println("zero")	
	}
	
}