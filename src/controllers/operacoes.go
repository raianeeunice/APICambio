package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/service"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// FazerDeposito é a função responsável
func FazerDeposito(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil{
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var deposito modelos.Deposito
	if erro = json.Unmarshal(corpoRequest, &deposito); erro != nil{
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = service.Preparar(&deposito); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	
	db, erro := banco.Conectar()
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeDepositos(db)
	deposito.ID, erro = repositorio.CriarDeposito(deposito)
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	
	respostas.JSON(w, http.StatusCreated, deposito)
}

// BuscarDepositos é a função responsável por buscar os depósitos feitos no banco
func BuscarDepositos(w http.ResponseWriter, r *http.Request){
	moeda := strings.ToUpper(r.URL.Query().Get("moeda"))
	db, erro := banco.Conectar()
	if erro !=nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	repositorio := repositorios.NovoRepositorioDeDepositos(db)
	depositos, erro := repositorio.BuscarDepositos(moeda)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusOK, depositos)
}

// BuscarSaldoTotal é a função responsável por buscar todo o saldo do banco
func BuscarSaldoTotal(w http.ResponseWriter, r *http.Request){
	moeda := strings.ToUpper(r.URL.Query().Get("moeda"))
	db, erro := banco.Conectar()
	if erro !=nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	repositorio := repositorios.NovoRepositorioDeDepositos(db)
	valorTotal, erro := repositorio.BuscarSaldoTotal(moeda)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusOK, valorTotal)
}
