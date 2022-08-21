package main

import "fmt"

type Address struct {
	neighborhood string
	number       uint32
}

type User struct {
	name    string
	age     uint8
	address Address
}

func main() {
	fmt.Println("Arquivo struct")

	var u User
	u.name = "Carlos"
	u.age = 120

	fmt.Println(u)

	user2 := User{"Carlos", 109, Address{"Dos Bobos", 53}}
	fmt.Println(user2)

	user3 := User{name: "Forduncio"}
	fmt.Println(user3)
}
