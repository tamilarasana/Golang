package main

import "fmt"

func main()  {

	slice1 := []int{1,2,3}
	slice2:= make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
	
}//sllice 

// func main()  {
// 	var a[]int  =  []int{1,4,5,6,3,5,7,3,2,7,8,9,10,11,12}
// 	var2 := append(a, 4,5)	
// 	for i:=0; i<len(a); i++{
// 		fmt.Println(a[i], var2)
// 	}
	
// }
 //ranege

/*func main()  {
	var a[]int = []int{1,4,5,6,3,5,7,3,2,7,8,9}
	for i, element := range a {
		fmt.Printf("%d: %d \n", i,element)

		


		
	}
	
}*/

/*func main()  {
	 var a[]int = []int{1,4,5,6,3,5,7,3,2,7,8,9}
	 for i, element := range  a{
		 for j,  element2 := range a{
			 if element == element2 && i != j {
				fmt.Println(element)
			 }
		 }
		 
	 }
	
}*/


/*func main()  {
	var a[]int = []int{1,4,5,6,3,5,7,3,2,7,8,9}
	for i, element := range  a{
		for j := i+1; j < len(a); j++ {
			element2 := a[j]
			if element2 == element{
				fmt.Println(element)
			}
		}
	}
	
	
}*/


