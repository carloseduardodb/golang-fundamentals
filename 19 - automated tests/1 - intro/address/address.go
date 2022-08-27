package address

import "strings"

func AddressType(address string) string {
	validTypes := []string{"Rua", "Avenida", "Praça", "Travessa", "Alameda", "Rodovia", "Estrada", "Lago", "Parque"}
	firstWord := strings.Split(address, " ")[0]
	isValid := false
	for _, validType := range validTypes {
		if strings.EqualFold(firstWord, validType) {
			isValid = true
			break
		}
	}
	if isValid {
		return firstWord
	}
	return "Invalid Type"
}
