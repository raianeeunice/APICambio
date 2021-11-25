package service

import (
	"api/src/modelos"
	"errors"
)

// Preparar é a função responsável de preparar o depósito para ser enviado ao BD
func Preparar(deposito *modelos.Deposito) error {
	if erro := validar(deposito); erro != nil {
		return erro
	}
	return nil
}

// validar é a função responsável por validar o campo ValorDepositado
func validar(deposito *modelos.Deposito) error {
	if deposito.ValorDeposito <= 0 {
		return errors.New("o valor depositado precisa se maior do que zero e não pode estar em branco")
	}
	return nil
}