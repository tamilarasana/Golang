package main

import  "fmt"

func main()  {
	age := 20
	

	if age >= 18 {
		fmt.Printf("you  can  raide!")
	}else{
		fmt.Println("No You Cannot raid")
		fmt.Printf("wait %d years", 18-age)
	}
	
}
