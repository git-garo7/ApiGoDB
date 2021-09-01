package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	//fmt.Println(config.StringConexaoBanco)//print for test.
	//fmt.Println("Running Api!") //print for test(go run main.go)

	r := router.Gerar() //call router and import.

	fmt.Printf("Lintening in door %d", config.Porta)                    //println(NOT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r)) //server.
}
