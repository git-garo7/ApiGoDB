package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//StringConexaoBanco=(string with connect database).
	StringConexaoBanco = ""
	//Porta=(port api running)
	Porta = 0
)

//Carregar=(initialize environment variable).
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}
	Porta, erro = strconv.Atoi(os.Getenv("API PORT!"))
	if erro != nil {
		Porta = 9000
	}
	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)
}

//.ENV=(variable enviroment guard)/got get github.com/joho/godottenv.
