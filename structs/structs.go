
package main

import "fmt"


//Structs
type car  struct{
	model string 
	topspeed int
	price int
	cc  int

}

func main()  {
	 var car1 car
	 car1. model = "tata"
	 car1.topspeed = 500
	 car1.price = 100000
	 car1.cc  	= 1200

	printCar(car1)
}

func printCar(car2 car)  {
	fmt.Println(car2.model)
	fmt.Println(car2.topspeed)
	fmt.Println(car2.price)
	fmt.Println(car2.cc)
}

/*func main()  {
	type car struct{
		model string
		topspeed int
		price int
	}
		var car1 car
		car1. model = "TATA"
		car1.topspeed = 500
		car1.price = 1000000
	fmt.Println(car1.model)
	fmt.Println(car1.topspeed)
	fmt.Println(car1.price)
}*/

/*type point struct{
	x int
	y int
}

type circle struct{
	radios float64
	center *point
}

func main()  {

	c1 := circle{2.5,&point{4,5}}
	fmt.Println("Structs program")
	fmt.Println(c1.center.x)
	
}*/


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

