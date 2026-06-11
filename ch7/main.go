package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func (u User) Greet() string {
	return fmt.Sprintf("Hello, %s!", u.Name)
}

func (u *User) Birthday() {
	u.Age++
}

const (
	A = iota + 10
	B
	C
)

func main() {

	user := User{Name: "Van", Age: 24}
	fmt.Println(user.Greet())

	user.Birthday() // equals to (&user).Birthday()

	fmt.Println(user.Age)

}
