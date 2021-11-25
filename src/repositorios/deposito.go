package repositorios

import (
	"api/src/modelos"
	"api/src/utils"
	"database/sql"
	"errors"
)

type Deposito struct {
	db *sql.DB
}

// NovoRepositorioDeDepositos cria um repositório de usuários
func NovoRepositorioDeDepositos(db *sql.DB) *Deposito {
	return &Deposito{db}
}

// CriarDeposito insere um depósito no Banco de Dados
func (d Deposito) CriarDeposito(deposito modelos.Deposito) (uint64, error) {
	statement, erro := d.db.Prepare(
		"INSERT INTO depositos (valorDeposito) VALUES (?)",
	)
	if erro != nil {
		return 0, nil
	}
	defer statement.Close()

	resultado, erro := statement.Exec(deposito.ValorDeposito)
	if erro != nil {
		return 0, nil
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, nil
	}
	return uint64(ultimoIDInserido), nil
}

// BuscarDepositos é função responsável por buscar todos os depositos
func (d Deposito) BuscarDepositos(moeda string) ([]modelos.Deposito, error) {
	qDepositos, erro := d.db.Query(
		"SELECT ID, valorDeposito, feitoEm FROM depositos",
	)
	if erro != nil {
		return nil, erro
	}
	defer qDepositos.Close()

	var depositos []modelos.Deposito
	for qDepositos.Next() {
		var deposito modelos.Deposito
		if erro = qDepositos.Scan(&deposito.ID, &deposito.ValorDeposito, &deposito.FeitoEm); erro != nil {
			return nil, errors.New("não há depositos em conta")
		}
		valorConvertido := utils.Converte(deposito.ValorDeposito, moeda)
		deposito.ValorDeposito = valorConvertido
		depositos = append(depositos, deposito)
	}

	return depositos, nil
}

// BuscarSaldoTotal é função responsável por buscar todo o saldo da conta em Real
func (d Deposito) BuscarSaldoTotal(moeda string) (modelos.Saldo, error) {
	depositado, erro := d.db.Query("SELECT SUM(valorDeposito) FROM depositos;")
	if erro != nil {
		return modelos.Saldo{}, erro
	}
	defer depositado.Close()

	var saldo modelos.Saldo
	
	depositado.Next()
	var soma float64
	if erro = depositado.Scan(&soma); erro != nil {
		return modelos.Saldo{}, errors.New("não há saldo em conta")
	}
	valorConvertido := utils.Converte(soma, moeda)
	saldo.ValorTotal = valorConvertido

	return saldo, nil
}
