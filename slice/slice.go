package main

import "fmt"

//sllice 

func main()  {
	var a[]int =  []int{1,4,5,6,3,5,7,3,2,7,8,9}
	for i:=0; i<len(a); i++{
		fmt.Println(a[i])
	}
	
}
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


