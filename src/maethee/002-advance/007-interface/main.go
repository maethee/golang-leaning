package main

import "fmt"

func main() {
	SampleMethodAndInterface()
}

type Circle struct {
	radius float64
}

func (c *Circle) area() float64 {
	fmt.Println("Area")
	fmt.Printf("TYPE of circle is %T\n", c)
	fmt.Printf("value of circle is %v\n  ", c)
	return c.radius
}

func (c *Circle) other() float64 {
	fmt.Println("Other")
	fmt.Printf("TYPE of circle is %T\n", c)
	fmt.Printf("value of circle is %v\n  ", c)
	return c.radius + 33
}

type Shape interface {
	area() float64
	other() float64
}

// functional programing should have interface for make sure who that using this func will have shape or data that function need.
func info(s Shape) string {
	return fmt.Sprintln("Shape is", s.area(), s.other())
}

// *Circle is POINTER can used just POINTER like &c
// Circle is a NON POINTER can use with c and &c
func SampleMethodAndInterface() {

	c := Circle{10}
	b := Circle{12}
	i := info(&c)
	fmt.Println("Info ", i)
	fmt.Println("Info ", b)
}
