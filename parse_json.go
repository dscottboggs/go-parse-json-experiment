package main

/*
Based on the answer here

https://stackoverflow.com/questions/34015049/how-to-parse-nested-json-into-structs-in-go#34015356
*/

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type ChildStruct1 struct {
	Key3 string
	Key4 string
}
type ChildStruct2 struct {
	Key4 string
	Key5 string
}
type MainStruct struct {
	Key1         string
	Key2         string
	ChildStruct1 ChildStruct1
	ChildStruct2 ChildStruct2
}

func main() {
	f, err := os.Open("json_sample.json")
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()
	var data MainStruct
	err = json.NewDecoder(f).Decode(&data)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(data.ChildStruct1.Key3)
	fmt.Printf("%+v", data)
}
