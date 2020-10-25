package main

import "fmt"


/*type vertext  struct{
	x  int
	y  int
}
func main()  {
       fmt.Println(vertext{10,20})	
}*/

/*type point struct{
	x  int32
	y  int32
}
func main()  {
	var tamil point = point{24,25}
	var tamil2  point = point{32,45}
	fmt.Println(tamil.y)     //fmt.Println(tamil.y)
	fmt.Println(tamil2.y)
}*/

type point struct{
	x int
	y int
}

type circle struct{
	radios float64
	center *point
}

func main()  {

	c1 := circle{2.5,&point{4,5}}
	fmt.Println(c1.center.x)
	
}