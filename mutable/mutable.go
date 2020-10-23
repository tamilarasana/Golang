package main

import "fmt"

/*func main()  {
	var x []int = []int{3,5,6}
	y := x
	y [0] = 10
	fmt.Println(x,y)
}*/

//Using slice 
/*func  main()  {
	var  x map[string]int = map[string]int{"tamil":7}
	y := x
	y["sam"] = 8
	//x["7"] = 7
	fmt.Println(x,y)
}*/

//Using a  Array 
/*func main()  {
	var x [3]int = [3]int{2,4,5} 
	y := x
	y[0] = 100
	fmt.Println(x,y)
}*/

//immutable 

func changeFirst(slice []int)  {
	slice[0] = 100
	
}

func main()  {
	var x []int = []int{3,4,5,6}
	fmt.Println(x)
	changeFirst(x)
	fmt.Println(x)
}