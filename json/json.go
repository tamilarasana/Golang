package main

import ("fmt"
        "encoding/json"
		
)

func main()  {
	type Person struct{
		Name  string
		Age int
		Car []string
	}

	p1:=&Person {Name:"tamil", Age:24, Car:[]string{"tata","van"}}
	data,_ := json.Marshal(p1) 
	fmt.Println(string(data))

	data1 :=  '{"Name":"tamil", "Age":24, "Car":["tata","van"]}'
	p2 := &Person{}
	json.Unmarshal([]byte(data1), p2)
	fmt.Println(p2.Name)
}