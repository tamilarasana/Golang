package main

import "fmt"

func main ()  {
	for x:=0; x<=10; x++{
		if x !=0 && x%2 ==0 {     //Only divided by  2 
			fmt.Println(x)
			continue
		}
	}
	
}

/*func main ()  {
	for x:=0; x<=20; x++{
		if x !=0 && x%2 ==0  && x%5 ==0 {     //Only divided by  2  and 5
			fmt.Println(x)
			continue
		}
	}
	
}*/
