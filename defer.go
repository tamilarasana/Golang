package main

import "fmt"

// func main()  {
// 	defer printText("defered")
// 	printText("tamil")
	
// }

func main()  {
	defer printText("first")
	printText("Tamilarasan") 
	defer printText("second")
	defer printText("third")
	defer printText("four")
	
}
func printText(p string)  {
	fmt.Println(p)
	
}