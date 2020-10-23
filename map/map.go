package main

import "fmt"

func main()  {
	 var mp map[string]int = map[string]int {

		 "tamil" : 7,
		 "vicky" : 8,
		 "ram"   : 9,
		 "sam"   : 5,
	 }
	 fmt.Println(mp["tamil"])
	 mp["tamil"] = 10    //update
	 delete(mp, "sam")   //delete
     mp["dude"] = 5     //Add new value 
	 fmt.Println(mp)

	
}