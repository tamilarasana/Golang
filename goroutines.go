package main

import ( "fmt"
		 "time"
)
func main()  {
	 go myfunc("first")
	 go myfunc("second")
	 go myfunc("third")

	 fmt.Println("done")
	 time.Sleep(time.Second)
	 fmt.Println("done")
	
}

func myfunc(a string)  {

	for i := 0; i < 3; i++ {
		fmt.Println(a, ":",i)
	}
	
}