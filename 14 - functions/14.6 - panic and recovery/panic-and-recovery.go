package main

import "fmt"

func recoveryExecution() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func studentApproved(n1, n2 float64) bool {
	defer recoveryExecution()
	average := (n1 + n2) / 2
	if average > 6 {
		return true
	} else if average < 6 {
		return false
	}
	panic("A média é exatamente 6")
}

func main() {
	fmt.Println(studentApproved(6.0, 6.0))
	fmt.Println("Pós execução!")
}
