package address_test

import (
	. "introduction-to-tests/address"
	"testing"
)

type mockTestCase struct {
	address  string
	expected string
}

func TestAddressType(t *testing.T) {
	mock := []mockTestCase{
		{"Rua das borboletas", "Rua"},
		{"Avenida das borboletas", "Avenida"},
		{"Praça das borboletas", "Praça"},
		{"Travessa das borboletas", "Travessa"},
		{"Alameda das borboletas", "Alameda"},
		{"Rodovia das borboletas", "Rodovia"},
		{"Estrada das borboletas", "Estrada"},
		{"Lago das borboletas", "Lago"},
		{"Parque das borboletas", "Parque"},
		{"Invalid Type", "Invalid Type"},
	}
	for _, m := range mock {
		actual := AddressType(m.address)
		if actual != m.expected {
			t.Errorf("AddressType(%s) = %s, expected %s", m.address, actual, m.expected)
		}
	}
}
