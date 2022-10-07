package main

import (
	"fmt"
)
//Triangle type
type triangle struct{
	height float64
	base float64
}
//Square type
type square struct{
	side float64
}

//Shape interface
type shape interface{
	getArea() float64
}

//printArea is a non-helper function
func printArea(s shape){
	fmt.Println(s.getArea())
}

//Square struct implementing getArea of shape interface
func (sq square) getArea() float64 {
	return sq.side * sq.side
}
//Triangle struct implementing getArea of shape interface
func (tr triangle) getArea() float64 {
	return (0.5 * tr.base * tr.height)
}

func main(){
	t := triangle{}
	s := square{}

	t.base = 10
	t.height = 10

	s.side = 4

	//Print area of square
	printArea(s)
	//Print area of triangle
	printArea(t)


}
