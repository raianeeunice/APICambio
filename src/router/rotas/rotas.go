package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa a estrutura de todas as rotas da API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
}

// Configurar é uma função que retorna o router já com todas as rotas configuradas
func Configurar(r *mux.Router) *mux.Router{
	rotas := rotaOperacoes

	for _, rota := range rotas{
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}