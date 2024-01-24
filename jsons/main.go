package main

import (
	"encoding/json"
	"fmt"
)

// declered struct
type myStruct struct {
	Name   string
	Age    int
	Status bool
	Values []int
}

func main() {
	// create struct
	s := myStruct{
		Name:   "John Connor",
		Age:    35,
		Status: true,
		Values: []int{15, 11, 37},
	}

	// struct to json
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n\n", data)

	// json to struct
	var NewStruct myStruct
	if err := json.Unmarshal(data, &NewStruct); err != nil {
		panic(err)
	}
	fmt.Printf("%v\n\n",NewStruct)

	// json to struct with tab
	data, err = json.MarshalIndent(NewStruct,  "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

}
