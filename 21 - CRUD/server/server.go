package server

import (
	"crud/bd"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	Id    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CRUD

// CREATE - POST
// Cria um novo usuário no banco de dados
func CreateUser(w http.ResponseWriter, r *http.Request) {
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
	if err = json.Unmarshal(requestBody, &user); err != nil {
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
	statement, err := db.Prepare("insert into usuarios (nome, email) values ($1, $2) RETURNING id")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao criar o statement"))
		return
	}

	var id int64
	err = statement.QueryRow(user.Nome, user.Email).Scan(&id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao obter o ID do último item inserido"))
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(fmt.Sprintf("Usuário criado com sucesso\nID: %d", id)))
}

// READ 01 - GET PER ID
// Busca um usuário específico
func SearchUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	id, err := strconv.ParseUint(parameters["id"], 10, 32)

	if err != nil {
		w.Write([]byte("Erro ao coletar o ID a ser pesquisado e convertê-lo para inteiro"))
		return
	}

	db, err := bd.ConectarBD()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}

	defer db.Close()

	line, err := db.Query("select * from usuarios where id = $1", id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao Buscar o Usuário"))
		return
	}

	defer line.Close()

	var user usuario
	if line.Next() {
		if err := line.Scan(&user.Id, &user.Nome, &user.Email); err != nil {
			w.Write([]byte("Erro ao Escanear o Usuário"))
			return
		}
	}

	w.WriteHeader(http.StatusAccepted)
	if err := json.NewEncoder(w).Encode(user); err !=nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao criar corpo do usuário em formato JSON"))
		return
	}
}

// READ 02 - GET THEM ALL
// Busca todos os usuários do banco de dados
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	db, err := bd.ConectarBD()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}

	defer db.Close()

	lines, err := db.Query("select * from usuarios")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao Buscar os Usuários"))
		return
	}

	defer lines.Close()

	var users []usuario

	for lines.Next(){
		var user usuario

		if err := lines.Scan(&user.Id, &user.Nome, &user.Email); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Erro ao criar corpo de usuários"))
			return
		}

		users = append(users, user)
	}

	w.WriteHeader(http.StatusAccepted)
	if err := json.NewEncoder(w).Encode(users); err !=nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao criar corpo de usuários em formato JSON"))
		return
	}
}

// UPDATE - PUT INSIDE NEW TOYS AND TAKE OUT THE BROKEN ONES
// Atualiza um usuário específico no banco de dados
func UpdateUser(w http.ResponseWriter, r *http.Request)  {
	parameters := mux.Vars(r)

	id, err := strconv.ParseUint(parameters["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao converter paramentro de ID para inteiro"))
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição"))
		return
	}

	var user usuario
	if err := json.Unmarshal(requestBody, &user); err != nil{
		w.Write([]byte("Erro ao converter usuário para struct"))
		return
	}

	db, err := bd.ConectarBD()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao se conectar ao banco de dados"))
		return
	}

	defer db.Close()

	statement, err := db.Prepare("update usuarios set nome = $1, email = $2 where id = $3")
	if err != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}

	defer statement.Close()

	if _, err := statement.Exec(user.Nome, user.Email, id); err != nil{
		w.Write([]byte("Erro ao atualizar o usuário"))
		return
	}

	w.Write([]byte("Usuário atualizado com sucesso"))
}

// DELETE - DELETE AND KILL THAT MOT********R
// Deleta um usuário específico no banco de dados
func DeleteUser(w http.ResponseWriter, r *http.Request)  {
	parameters := mux.Vars(r)

	id, err := strconv.ParseUint(parameters["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao converter paramentro de ID para inteiro"))
		return
	}

	db, err := bd.ConectarBD()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao se conectar ao banco de dados"))
		return
	}

	defer db.Close()

	statement, err := db.Prepare("delete from usuarios where id = $1")
	if err != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}

	defer statement.Close()

	if _, err := statement.Exec(id); err != nil{
		w.Write([]byte("Erro ao deletar o usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Usuário deletado com sucesso"))
}