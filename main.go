package main

import (
	"fmt"
	"net/http"
	"projeto/database"
	"projeto/handlers"
)

func main() {
	// Conectar ao banco
	db, err := database.InitDB()
	if err != nil {
		fmt.Println("Erro ao conectar:", err)
		return
	}
	defer db.Close()
	fmt.Println("Conectado ao banco com sucesso!")

	// Rotas
	http.HandleFunc("/usuario", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.ListarUsuarios(w, r)
		} else if r.Method == http.MethodPost {
			handlers.CriarUsuario(w, r)
		}
	})

	http.HandleFunc("/usuario/", handlers.BuscarUsuario)

	fmt.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
