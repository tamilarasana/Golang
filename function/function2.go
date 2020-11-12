package main

import "fmt"

func main(){
	tamil := []int{10,20,30,40,50}
	total := 0
	 for  _, as := range tamil{

	 	total += as

	 }
	 fmt.Println(total/len(tamil))
}
