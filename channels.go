package main

import "fmt"

func main()  {
	myChannels := make(chan string)
	go func (){
		 myChannels<- "A Tamilarasan"
		 myChannels<- "Annamalai" 
	         }()

	text := <-myChannels
	//fmt.Println(text, <-myChannels)
	fmt.Println(text)
	fmt.Println(<-myChannels)
		
	}
