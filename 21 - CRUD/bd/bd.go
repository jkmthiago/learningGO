package bd

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConectarBD() (*sql.DB, error){
	user 		:= os.Getenv("DB_USER")
	host 		:= os.Getenv("DB_HOST")
	port 		:= os.Getenv("DB_PORT")
	password 	:= os.Getenv("DB_PASS")
	dbName 		:= os.Getenv("DB_NAME")

	var dataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	fmt.Println("Iniciando comunicação com o banco de dados")

	db, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		fmt.Println("Conexão Falha com o Banco: Erro de Inicio de Conexão")
		return nil, err
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Conexão Falha com o Banco: Erro de Ping")
		return nil, err
	}

	fmt.Println("Conexão Bem Sucedida com o Banco:")
	return db, nil
}
