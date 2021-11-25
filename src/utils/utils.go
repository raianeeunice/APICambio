package utils

import (
	"log"
)

// Converter é a função responsavel por converter os valores de Real para as moedas especificadas
func Converte(valor float64, moeda string) float64 {
	switch moeda {
	case "EUR": // euro
		return valor / calculaCambio(6.25)
	case "GBP": // libra
		return valor / calculaCambio(7.33)
	case "USD": // dolar
		return valor / calculaCambio(5.46)
	case "JPY": // iene
		return valor / calculaCambio(0.048)
	default:
		return valor
	}
}

// calculaCambio é a função que vai calcular o cambio com as taxas IOF, Taxa de cambio e retorna o valorFinal
func calculaCambio(valorMoeda float64) float64 {
	if valorMoeda <= 0 {
		log.Fatal("o valor da moeda não pode ser menor ou igual a zero")
	}
	iof := 0.0638
	taxaCambio := 0.16
	valorFinal := valorMoeda + (valorMoeda * iof) + (valorMoeda * taxaCambio)
	return valorFinal
}