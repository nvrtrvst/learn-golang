package main

import "fmt"

func main() {
	var m map[string]int
	m = map[string]int{}


	m["Answer"] = 42
	fmt.Println("The Value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The Value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The Value:", m["Answer"])

	fmt.Println("The Value:", m["Answer"])

	if v, ok := m["ada"]; ok {
	fmt.Println("The Value:", v, m,["Present?" ok])
	}
}
	





