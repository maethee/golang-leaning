package main

import "fmt"

func main() {
	SampleStruct()
}

//struct is a aggregate data types
// x := type{values} composite literal
type person struct {
	name     string
	lastname string
}

type secretPerson struct {
	person //undefine type and name
	allow  bool
}

type GroupType struct {
	name, lastname string
	age, money     int
}

func SampleStruct() {
	fmt.Println("SampleStrut")
	p1 := person{
		name:     "Maethee",
		lastname: "Cha",
	}

	//short style the value assign in sequence struct (name, lastname)
	p2 := person{"Bob", "Lee"}

	fmt.Println(p1)
	fmt.Println(p1.name, p1.lastname)

	fmt.Println(p2)
	fmt.Println(p2.name, p2.lastname)

	seP1 := secretPerson{
		person: person{
			name:     "Jame",
			lastname: "Bond", // should trailing comma every line
		},
		allow: true,
	}

	fmt.Println(seP1)
	// person.name and name is the same value when top leven not have like name it will promote automaticly from outer TYPE secretPerson
	fmt.Println(seP1.person.name, seP1.name)

	//multiple assigning
	//we can assign values at one line
	x := 1
	y := 2
	fmt.Println(x, y)
	x, y = y, x

	fmt.Println(x, y)
	x, y = 1, 2

	//anonymous structs
	//sometime you want to using strcut for small thing want code to make it clean outsite
	ano := struct {
		name string
		dest string
	}{
		name: "secret",
		dest: "mission",
	}

	fmt.Println(ano)
}
