package utils

import "testing"

func TestConverte(t *testing.T) {
	esperado := Converte(2762, "JPY")
	if esperado != 47018.84839570736 {
		t.Errorf("Função esperava: 47018.84839570736, retornou: %f", esperado)
	}

	esperado2 := Converte(2762, "USD")
	if esperado != 413.35251336885597 {
		t.Errorf("Função esperava: 413.35251336885597, retornou: %f", esperado2)
	}
}
