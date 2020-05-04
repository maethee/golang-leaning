package main

import "fmt"

func main() {
	SampleMap()
}

func SampleMap() {
	fmt.Println("SampleMap")

	//Map is key-value
	m := map[string]int{
		"Maethee": 33,
		"Bee":     34,
	}
	fmt.Println(m)
	fmt.Println(m["Maethee"])
	fmt.Println(m["xxx"])

	//we can ask the value has in a map like below

	v, ok := m["test"]
	fmt.Println(v)
	fmt.Println(ok)

	//check in condition
	if v, ok := m["Maethee"]; ok {
		fmt.Println("Maethee", v)
	}

	//add element to map
	m["new people"] = 55
	fmt.Println(m)

	//delete map
	delete(m, "Bee")
	fmt.Println(m)
}
