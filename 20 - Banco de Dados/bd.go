package main

import (
	"database/sql"
	"fmt"
	"log"
	// "os"
	"flag"

	_ "github.com/lib/pq"
)

func main() {
	// leitura e captura por arquivo yml

	// user := os.Getenv("DB_USER")
	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	// password := os.Getenv("DB_PASS")
	// dbName := os.Getenv("DB_NAME")

	user 	 := flag.String("u", "", "Insere o nome do usuário")
	host 	 := flag.String("h", "", "Insere o host de acesso")
	port 	 := flag.String("port", "", "Insere a porta de acesso")
	password := flag.String("pass", "", "Insere a senha do usuário")
	dbName 	 := flag.String("dbName", "", "Insere o nome do banco a ser acessado")

	flag.Parse()

	if *user == "" || *host == "" || *port == "" || *password == "" || *dbName == ""{
		fmt.Println(*user, *host, *port, *password, *dbName)
		fmt.Println("Favor inserir corretamente todos os parametros de acesso")
		fmt.Println("-u ----> Usuário")
		fmt.Println("-h ----> Host")
		fmt.Println("-port -> Porta de Acesso")
		fmt.Println("-pass -> Senha")
		fmt.Println("-dbName ---> Banco de Dados")

		log.Fatal()
	}

	var dataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", *host, *port, *user, *password, *dbName)
	
	fmt.Println(*host, *port, *user, *password, *dbName)

	db, err := sql.Open("postgres", dataSourceName)
	
	if err != nil {
		fmt.Println("Conexão Falha com o Banco: Erro de Inicio de Conexão")
		log.Fatal(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Println("Conexão Falha com o Banco: Erro de Ping")
		log.Fatal(err)
	}

	fmt.Println("Conexão Bem Sucedida com o Banco:", db)

	lines, err := db.Query("select * from usuarios")

	if err != nil {
		log.Fatal(err)
	}

	defer lines.Close()

	fmt.Println(lines)
}
