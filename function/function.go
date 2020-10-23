package main

import  "fmt"

func max(x int ,y int )int  {
	if x > y {
	  return x
	}else{
		return y
	}
}

func main()  {
	x := max(10,6)
	fmt.Println(x)
	
}

/*func add(x,y int) (int, int)  {
	return x + y, x - y
	
}
func main()  {

	ans1, ans2  := add(10, 7)
	fmt.Println(ans1,ans2)
}*/

/*func add(x,y int)(z1 int,  z2 int)  {
	       z1 = x+y 
		   z2= x-y
		   return
	
}

func main()  {
	ans1,ans2 := add(20,30)
	fmt.Println(ans1,ans2)
	
}*/


//advance function  concepts and Closurs

// func main()  {
// 	test := func ()  {
// 		fmt.Println("hello")
		
// 	}
// 	test()
	
// }

// func main()  {
// 	test := func (x int) int  {
// 		return x*2		
// 	 }
// 	test(2)

// 	fmt.Println(test)
// }

/*func test1(myFunc func(int)int)  {
	fmt.Println(myFunc(7))
	
}

func main()  {
	test := func (x int) int  {
		return x*7
		
	} 
	test1(test)
}*/
/*func returnFunc(x string) func ()  {
	sum := 2
	return func ()  {
		fmt.Println(sum)
		
	}
	
}  
	
func main()  {
	returnFunc("helllo")()
	x := returnFunc("goodbye")
	x()
	
}*/


