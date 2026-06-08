package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz: ", i)
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz: ", i)
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz: ", i)
			continue
		}
		fmt.Println(i)
	}
	evenVals0 := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals0 {
		fmt.Println(i, v)
	}

	evenVals1 := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals1 {
		fmt.Println(v)
	}

	m := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}
	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}
	///////////////////////////////////////////////////////////////////////
	words := []string{"a", "cow", "smile", "gopher"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}

	for i := 0; i < 10; i++ {
		switch i {
		case 0, 2, 4, 6:
			fmt.Println(i, "is even")
		case 3:
			fmt.Println(i, "is divisible by 3 but not 2")
		case 7:
			fmt.Println("exit the loop!")
			break
		default:
			fmt.Println(i, "is boring")
		}
	}

}
