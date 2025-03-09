package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	//go get github.com/go-sql-driver/mysql
	"github.com/joho/godotenv"
	//go get github.com/joho/godotenv
)

var DB *sql.DB

func ConnectDB() {
	// Carregar variáveis do .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	// Criar a string de conexão
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Conectar ao MySQL
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	// Testar conexão
	err = DB.Ping()
	if err != nil {
		log.Fatal("Erro ao verificar conexão com o banco:", err)
	}

	fmt.Println("📌 Conectado ao banco de dados com sucesso!")
}
