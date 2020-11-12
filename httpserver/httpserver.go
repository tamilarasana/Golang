package main

import ( "fmt"
		 "net/http"
)
// func main()  {
// 	http.HandleFunc ("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w,"hello Tamilarasan")
	
// 	})

// 	http.HandleFunc ("/tamil", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w,"hello frinds")
	
// 	})

// 	http.ListenAndServe(":9000",nil)
	
// }


//Another way 

func main()  {
	myHandlerfunc := func (w http . ResponseWriter, r* http.Request)  {
		fmt.Fprintf(w, "hello frinds")

		
	}

	http.HandleFunc("/", myHandlerfunc)
	http.HandleFunc("/name", myHandlerfunc)

	http.ListenAndServe(":9000", nil)
	
}

func myNameHandleerfunc(w  http.ResponseWriter, r* http.Request)  {
	fmt.Fprintf(w, "Hello Tamilarasan")
	
}
