package main

import ( "fmt"
		"os"
		"bufio"
		  "io/ioutil"
)

func main()  {
	massage :=  []byte("Hello Tamilarasan")
	err :=  ioutil.WriteFile("new.txt",massage,0644)
	if err != nil{
		fmt.Println(err)
	}

		//Creating new file
	f,err := os.Create("newtext.txt")
	defer f.Close()
	if err != nil{
		fmt.Println(err)
	}
	f.WriteString("Hello Frinds")
	f.Sync()

	//writing a new file using bufio

	w := bufio.NewWriter(f)

	w.WriteString("How are you ?")
	w.Flush()

	
} 
