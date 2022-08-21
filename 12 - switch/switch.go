package main

import "fmt"

func dayWeekReturn(number int) string {
	switch number {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Numero inválido"
	}
}

func dayWeekReturnInternalCondition(number int) string {
	var dayWeek string
	switch {
	case number == 1:
		dayWeek = "Domingo"
		fallthrough
	case number == 2:
		dayWeek = "Segunda-Feira"
	case number == 3:
		dayWeek = "Terça-Feira"
	case number == 4:
		dayWeek = "Quarta-Feira"
	case number == 5:
		dayWeek = "Quinta-Feira"
	case number == 6:
		dayWeek = "Sexta-Feira"
	case number == 7:
		dayWeek = "Sábado"
	default:
		dayWeek = "Numero inválido"
	}
	return dayWeek
}

func main() {
	fmt.Println("Switch")

	dayOne := dayWeekReturn(3)

	fmt.Println(dayOne)

	dayTwo := dayWeekReturnInternalCondition(5)

	fmt.Println(dayTwo)
}
