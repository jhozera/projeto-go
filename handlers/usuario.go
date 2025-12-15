package handlers

import (
	"encoding/json"
	"net/http"
	"projeto/models"
	"projeto/services"
	"strconv"
)

var usuarioService = &services.UsuarioService{}

// GET /usuario
func ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	usuarios := usuarioService.ListarTodos()
	json.NewEncoder(w).Encode(usuarios)
}

// GET /usuario/:id
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	idInt, _ := strconv.Atoi(id)
	usuario := usuarioService.BuscarPorID(idInt)

	if usuario == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Message: "Usuário não encontrado",
			Code:    404,
		})
		return
	}

	json.NewEncoder(w).Encode(usuario)
}

// POST /usuario
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.CriacaoUsuario
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Message: "Requisição inválida",
			Code:    400,
		})
		return
	}

	usuario := usuarioService.CriarUsuario(req.Nome, req.Email, req.Senha)
	if usuario == nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Message: "Erro ao criar usuário",
			Code:    500,
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(usuario)
}
