package main

import  "fmt"

func main()  {

    //AND 
	val1 := true && true    //true
	val2:= true  && false   //false
	val3:= false && true 	//false
	val4:= false && false	//false

	//OR
 	val5 := true  || true   //true 
	val6 := true  || false 	//true
	val7 := false || true 	//true
	val8:= false  || false  //false

	//NOT
	val9 :=    !true 		//false
	val10 :=   !false 		//true
	x := 5

	var1 :=  (true || false) && true   //true
	var2 := false || true  && true     //true
	var3 := false || false && false    //false
	var4 := (true || false) && !false  //true\
	var5 := (true || false) && !false  || x < 10  //true

	

	fmt.Printf("%t \n", val1)
	fmt.Printf("%t \n", val2)
	fmt.Printf("%t \n", val3)
	fmt.Printf("%t \n", val4)
	fmt.Printf("%t \n", val5)
	fmt.Printf("%t \n", val6)
	fmt.Printf("%t \n", val7)
	fmt.Printf("%t \n", val8)
	fmt.Printf("%t \n", val9)
	fmt.Printf("%t \n", val10)

	fmt.Printf("%t \n", var1)
	fmt.Printf("%t \n", var2)
	fmt.Printf("%t \n", var3)
	fmt.Printf("%t \n", var4)
    fmt.Printf(" %t \n", var5)
	
	
}