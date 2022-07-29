package main

import (
	"encoding/json"
)

type Person struct {
	Name   string
	Age    int64
	Weight float64
}

func main() {
	//person := Person{
	//	Name:   "Wang Wu",
	//	Age:    30,
	//	Weight: 150,
	//}
	//jsonBytes, _ := json.Marshal(person)
	//fmt.Println(string(jsonBytes))
	str1 :=  `{"Name":"Wang Wu","Age":30,"Weight":150}`
	var personFromJSON interface{}
	json.Unmarshal(str1, &personFromJSON)

	//r := []type(str).(map[string]interface{})
	//fmt.Println(r)
	//fmt.Println(reflect.TypeOf(r["Age"]).Name())  // float64
	//n := r["Weight"].(int)
	//fmt.Println(n)
	//fmt.Println(n, ok)
}
