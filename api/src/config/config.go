package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//String de conexão com o mysql
	StringConexaoBanco = ""

	//Porta onde a aAPI estará rodando
	Porta = 0

	//SecretKey é a chave utilizada pra assinar o Token
	SecretKey []byte
)

// Carregar irá iniciar as variáveis de ambiente
func Carregar() {

	var erro error
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)
	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
