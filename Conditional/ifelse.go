


package main

import "fmt"


func main() { 
	var mark int
	fmt.Printf("Enter your mark : ")
	fmt.Scanf("%d", &mark)
	fmt.Println("Your Mark : ", mark)
	fmt.Printf("Your Result : ")

	if  mark >= 90 && mark <= 100 {
		fmt.Println("First class")
	}else if mark >= 75 && mark < 90 {
		fmt.Println("Second  Class")
	}else if mark >=50 && mark < 75 {
		fmt.Println("Third class")
	}else if mark >=35 && mark < 50 {
		fmt.Println("Your Pass")
	}else {
		fmt.Println("Your fail")
	}
}

// func main()  {
// 	age := 15
// 	if age >= 18{
// 		fmt.Printf("You can raide")
// 	}else if  age >= 14{
// 		fmt.Printf("You can ride with your parents")
// 	}else {
// 		fmt.Printf("you cannot raide")
// 	}
	
// }



