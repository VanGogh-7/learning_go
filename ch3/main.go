package main

import (
	"fmt"
	"slices"
)

func main() {
	var x [3]int
	x[0] = 10

	var y = [12]int{1, 5: 4, 6, 10: 100, 15}
	y[10] = 13
	fmt.Println(x, len(x), cap(x))
	var z = [...]int{1, 2, 3}

	fmt.Println(z)
	var x1 = []int{1, 2, 3}
	x1 = append(x1, 4)

	fmt.Println(x1)
	x2 := []int{1, 2, 3, 4, 5}
	y1 := []int{1, 2, 3, 4, 5}
	z1 := []int{1, 2, 3, 4, 5, 6}
	//s := []string{"a", "b", "c"}
	fmt.Println(slices.Equal(x2, y1))
	fmt.Println(slices.Equal(x2, z1))
	//fmt.Println(slices.Equal(x2, s))
	x3 := make([]int, 5)
	x3 = append(x3, 4)
	fmt.Println(x3)
	clear(x3)
	fmt.Println(x3)

	////////////////////////////////////////////////////

	xSlice := []int{1, 2, 3}
	xArray := [3]int(xSlice)
	smallArray := [2]int(xSlice)
	xSlice[0] = 10
	xArrayPointer := (*[3]int)(xSlice)

	fmt.Println(xSlice)
	fmt.Println(xArrayPointer)
	fmt.Println(smallArray)
	fmt.Println(xArray)
	/////////////////////////////////////////
	teams := map[string][]string{
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}
	fmt.Println(teams[teams["Orcas"][0]])
	fmt.Println(teams[teams["Orcas"][1]])
	teams["Van"] = []string{"ww", "qq", "ee"}

	////////////////////////////////////////
	type person struct {
		name string
		age  int
		pet  string
	}
	julia := person{"Julia", 40, "cat"}
	fmt.Println(julia)

}
