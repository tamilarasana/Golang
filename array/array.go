package main

import "fmt"

func main()  {
	arr:= [5]int{4,5,6,7,8}
	sum :=0

	arr2D :=[2][3]int{{1,2,3},{2,4,5}}
	fmt.Println(arr2D)

	for i:=0; i<len(arr); i++{
	sum  +=arr[i]
	}
	fmt.Println(sum)
	
}







