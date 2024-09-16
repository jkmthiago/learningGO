package server

import (
	"crud/bd"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type usuario struct{
	ID 		uint32 `json:"id"`
	NOME 	string `json:"nome"`
	EMAIL 	string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request)  {
	// CRIA O CORPO DA REQUISIÇÃO
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Falha na leitura do corpo da requisição"))
		return
	}

	// CRIA UM STRUCT VAZIO PARA ARMAZENAR S DADOS A SEREM ENVIADOS AO BANCO
	var user usuario

	// REALIZA O DESEMPACOTAMENTO DA REQUISIÇÃO E ARMAZENA OS VALORES NO STRUCT
	if err = json.Unmarshal(requestBody, &user); err != nil{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Falha na converção de requisição para struct"))
		return
	}

	// Cria conexão com o banco de dados
	db, err := bd.ConectarBD()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}

	defer db.Close()
	statement, err:= db.Prepare("insert into usuarios (nome, email) values ($1, $2)")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao criar o statement"))
		return
	}

	var id int32

	if err = statement.QueryRow(user.NOME, user.EMAIL).Scan(id); err!=nil {
		w.WriteHeader(http.StatusAccepted + http.StatusInternalServerError)
		w.Write([]byte("Usuário criado mas não foi possível obter seu ID"))
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(fmt.Sprintf("Usuário criado com sucesso\nID: %d", id)))
}