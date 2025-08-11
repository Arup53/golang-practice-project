package main

import "fmt"

// why receiver function needs to be pointer receiver in go
// from oops sense, we are directly  modifing the properties of a object so, in go we use pointer receiver or return pointer to modify the struct(properties)

type Koi struct {
	Name   string
	Origin string
}

type Shapes struct {
	Name string
}

type Pond struct {
	Name string
	Shapes ShapesList
	Fish []Koi
}

type KoiList []Koi

type ShapesList []Shapes

func main() {

	// koiArrays := []Koi{
	// 	{Name: "blue",
	// 		Origin: "osaka",
	// 	}, {
	// 		Name:   "red",
	// 		Origin: "tokyo",
	// 	},
	// }

	// koiArrays2 := KoiList{
	// 	{Name: "red", Origin: "shifu"},
	// 	{Name: "orange", Origin: "yokhoma"},
	// }

	// pond := Pond{
	// 	Name: "Green",
	// 	Shapes: ShapesList{
	// 		{Name: "square"},
	// 		{Name: "round"},
	// 	},
	// 	Fish: []Koi{
	// 		{Name: "blue",Origin: "osaka",}, 
	// 		{Name:"red",Origin:"tokyo",},
	// 	},
	// }
     
    
	// fmt.Println(koiArrays2)
	x:=2
	px:= &x
	fmt.Println(*px)
}