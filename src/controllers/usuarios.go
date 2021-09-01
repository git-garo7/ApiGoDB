package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//CriarUsuario=(insert user into database)
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}
	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarioID, erro := repositorio.Criar(usuario)
	if erro != nil {
		log.Fatal(erro)
	}
	w.Write([]byte(fmt.Sprintf("Id inserido: %d", usuarioID)))
}

//(BuscarUsuarios=(search  all users in database)
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("search users!"))
}

//BuscarUsuario=(search user in database)
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("search user!"))
}

//AtualizarUsuario=(update data user in database)
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("updating user!"))
}

//DeletarUsuario=(delete data user)
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("deleting user!"))
}
