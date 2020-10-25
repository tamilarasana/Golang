package main

import "fmt"



func main(){
	var i int = 5
	var ptr  *int
	var pptr **int

	ptr = &i
	pptr  = &ptr
	//fmt.Println(i)
	fmt.Println(&ptr)
	fmt.Println(**pptr)
} 
// func main()  {
// 	tochange := "Tamilarasan"
// 	var pointer *string = &tochange
// 	fmt.Println(pointer) 
// 	fmt.Println(*pointer)
// 	fmt.Println(&pointer)
	
// }

// func main()  {
// 	x:= 7
// 	fmt.Println(&x)
	
// }

// func main()  {
// 	x:= 7
// 	y := &x
// 	fmt.Println(x,y)
// 	*y = 8
// 	fmt.Println(x,y)
	
// }

// func changeValue(str *string)  {
// 	*str = "Tamil"
	
// }

// func changeValue2(str string)  {
// 	str = "Changed"

	
// }
// func main()  {
// 	tochange := "Arasan"
// 	changeValue(&tochange)
// 	fmt.Println(tochange)
// 	changeValue2(tochange)
// 	fmt.Println(tochange)
	
// }




