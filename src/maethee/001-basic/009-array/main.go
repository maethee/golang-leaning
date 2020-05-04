package main

import "fmt"

func main() {
	SampleArray()
}

func SampleArray() {
	var x [5]int
	var y = [6]int{1, 2, 3, 4, 5, 6}
	// above x and y are different type

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))

	//slice
	//slice allows you to group together VALUES of the same TYPE
	//wrong: c := []int{4, 5, 6, "string", 8}
	//call comsosite literal
	c := []int{4, 5, 6, 7, 8}
	fmt.Println(c)
	fmt.Printf("%T\n", c)

	//each c contain int type
	for key, value := range c {
		fmt.Println(key, value)
		fmt.Printf("%T\t%T\n", key, value)
	}

	//slicing
	d := []int{4, 5, 6, 7, 8, 9}
	fmt.Println(d[0])
	fmt.Println(d[1:])  // give me from position 1 to the end  = [5 6 7 8 9]
	fmt.Println(d[1:4]) // give me from position 1 up to 4   = [5, 6, 7]

	//append the same TYPE
	d = append(d, 11, 22, 33, 44)
	fmt.Println(d)

	//
	e := []int{111, 222, 333}

	d = append(d, e...) //e... pull whole bunch values and put in to it "d"
	fmt.Println(d)
	//notes ...T can receive unlimit parameters like func Unlimit(...interface{}) variadic parameter  // we can receive unlimit interface that push to this func
	//Unlimit("sdf", 44, 55)

	//slice deleting
	d = append(d[:2], d[4:]...)
	fmt.Println(d)

	//make type,len,size
	f := make([]int, 10, 10)
	fmt.Println(f)
	fmt.Println(len(f))
	fmt.Println(cap(f))

	//x[10] = we can't assign in position 10 because we declare just 10 (0 - 9)
	//but we can used append because append will create a new one and type to hold the data.

	f = append(f, 2344)
	fmt.Println(f)
	fmt.Println(len(f))
	fmt.Println(cap(f))

	f = append(f, 2345) //append is a copy array to new array
	fmt.Println(f)
	fmt.Println(len(f))
	fmt.Println(cap(f))
	// 0 0 0 0 0 0 0 0 0 0 2344 2345]
	// append will automatic change underlying array to hole all of that data

	// slice string and multiple dimensional
	g := []string{"maethee", "foo", "bar"}
	fmt.Println(g)

	h := []string{"Mr", "Ms", "Mrs"}
	fmt.Println(h)

	dimension := [][]string{g, h}
	fmt.Println(dimension)
}
