package main  

import "fmt"
type person struct {
   name string
   age int 
}


func stactAsMapDataType () {

	a := map[string]person{
		"first" : {
			name : "pranto",
			age : 32,
		}  ,
		"second" : {
			name: "shovon",
			age : 23,
		},
		"third" : {
			name : "shamol",
			age : 20,
		},

	}
	fmt.Println(a)
}