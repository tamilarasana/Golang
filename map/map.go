package main

import "fmt"

			func main() {
				elements := map[string]map[string]int{
				"H": map[string]int{
						"name": 2,
						"age":1,
				 },
				 "He": map[string]int{
					"name":4,
					"age":4,
					},
				"T" : map[string]int{
					"name" : 5,
					"age" : 3,
				}
				"S" : map[string]int{
					"name" : 6,
					"age"  : 7
				}

				 
		    	}
				if el, ok := elements["H"]; ok {
					fmt.Println(el["name"], el["state"])
				}
			
			
		    }
			// 	"He": map[string]string{
			// 	"name":"Helium",
			// 	"state":"gas",
			// 	},
			// 	"Li": map[string]string{
			// 	"name":"Lithium",
			// 	"state":"solid",
			// 	},
		  
			//    "Be": map[string]string{
			// 	"name":"Beryllium",
			// 	"state":"solid",
			// 	},
			// 	"B": map[string]string{
			// 	"name":"Boron",
			// 	"state":"solid",
			// 	},
			// 	"C": map[string]string{
			// 	"name":"Carbon",
			// 	"state":"solid",
			// 	},
			// 	"N": map[string]string{
			// 	"name":"Nitrogen",
			// 	"state":"gas",
			// 	},
			// 	"O": map[string]string{
			// 	"name":"Oxygen",
			// 	"state":"gas",
			// 	},
			// 	"F": map[string]string{
			// 	"name":"Fluorine",
			// 	"state":"gas",
			// 	},
			// 	"Ne": map[string]string{
			// 	"name":"Neon",
			// 	"state":"gas",
			// 	},
			// 	}
			 	// if el, ok := elements["H"]; ok {
			 	// fmt.Println(el["name"], el["state"])
				// }
			// }
























// 	elements := make(map[string]string)
// 		elements["H"] = "Hydrogen"
// 		elements["He"] = "Helium"
// 		elements["Li"] = "Lithium"
// 		elements["Be"] = "Beryllium"
// 		elements["B"] = "Boron"
// 		elements["C"] = "Carbon"
// 		elements["N"] = "Nitrogen"
// 		elements["O"] = "Oxygen"
// 		elements["F"] = "Fluorine"
// 		elements["Ne"] = "Neon"
// 		fmt.Println(elements["Li"])
// 	}
// }

// 	x := []int{
// 		48,96,86,68,
// 		57,82,63,70,
// 		37,34,83,27,
// 		19,97, 9,17,
// 		}
// 		small := x[0]


// 	for _, element  :=  range x{
// 		if element < small{
// 			small =  element
// 		}
// 	} 
// 	fmt.Println("small",small)
// }

// 	var x[]int = make([]int, 3, 9)

// 	for  i :=0; i<len(x); i++{
// 		fmt.Println(x[i])
// 	}
// }
// 	fmt.Printf("ENTER HERE:")
// 	var ma string
//     fmt.Scanf("%s",&ma)
// 	elements := make(map[string]string)

// 		elements["T"] = "Tamiarasan"
// 		elements["V"] = "Vignesh"
// 		elements["S"] = "SAM"
// 		elements["R"] = "RAM"
// 		elements["M"] = "MANI"
// 		elements["K"] = "KANNAN"
// 		name, ok := elements[ma]

// 		 	if !ok {
// 		 		fmt.Println("Key is not Found")
// 			}
	
// 		fmt.Println(name)
// }	

// 	fmt.Printf("Enter the name :")
// 	var input string
// 	var mp map[string]int = map[string]int {

// 		"tamil" : 7,
// 		"vicky" : 8,
// 		"ram"   : 9,
// 		"sam"   : 5,
// 	}
	
// 	fmt.Scanf("%s",&input)
// 	name, ok := mp[input]

// 	if !ok {
// 		fmt.Println("Key is not Found")
// 	}
// 	fmt.Println("no of name: %d ", name)
// }


/*func main()  {
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

	
}*/