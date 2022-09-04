package main

import "fmt"

//type Supplement interface {
//Details() string
//}

type Supplement struct {
	Name   string
	Dosage int
	Unit   string
}

//Create empty map of Supplements and append supplements to it

func main() {
	var supplements []Supplement
	supplements = append(supplements, Supplement{"Vitamin C", 500, "mg"})
	supplements = append(supplements, Supplement{"Vitamin D", 1000, "IU"})

	for range supplements {
		fmt.Println(supplements)
	}
}
