package main

import "fmt"

type user struct {
	name string
	age  uint8
}

// atrela método dentro da estrutura do usuário
func (u user) save() {
	fmt.Println("Salvando usuário", u.name)
}

func (u user) ofAge() bool {
	return u.age >= 18
}

// é igual usar o this.u.age++ na POO
func (u *user) haveABirthday() {
	u.age++
}

func main() {
	// user1 := user{"João", 30}
	// user1.save()
	// fmt.Println(user1)
	user2 := user{"Maria", 18}
	fmt.Println(user2)
	fmt.Println(user2.ofAge())
	user2.haveABirthday()
	fmt.Println(user2)
}
