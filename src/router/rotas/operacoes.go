package rotas

import (
	"api/src/controllers"
	"net/http"
)

// rotaOperacoes é um slice de rotas que vai representar
// todas as rotas de operações do sistema
var rotaOperacoes = []Rota{
	{
		URI:    "/depositar",
		Metodo: http.MethodPost,
		Funcao: controllers.FazerDeposito,
	},

	{
		URI:    "/depositos",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarDepositos,
	},

	{
		URI:    "/saldoTotal",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarSaldoTotal,
	},
}