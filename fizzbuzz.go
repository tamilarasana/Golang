package main

import "fmt"

func main()  {

	fmt.Println("Fizz And Buzz Program")
	
	for num := 1; num<=50;  num++ {

		if num %15 == 0{
			fmt.Println("FizzBuzz")
		}else if num % 3 == 0{
			fmt.Println("Fizz")
		}else if  num % 5 == 0{
			fmt.Println("Buzz")
		}else{
			fmt.Println(num)
		}
		
	}
	
}