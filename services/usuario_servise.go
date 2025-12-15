package services

import (
	"fmt"
	"projeto/database"
	"projeto/models"

	"golang.org/x/crypto/bcrypt"
)

type UsuarioService struct {
}

func (s *UsuarioService) CriarUsuario(nome, email, senha string) *models.Usuario {
	fmt.Println("Senha recebida:", senha)

	// Criptografar a senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Erro ao criptografar:", err)
		return nil
	}

	fmt.Println("Hash gerado:", string(hashedPassword))

	// Salvar hash no banco
	result, err := database.DB.Exec("INSERT INTO usuarios (nome, email, senha) VALUES (?, ?, ?)", nome, email, string(hashedPassword))
	if err != nil {
		fmt.Println("Erro ao inserir:", err)
		return nil
	}

	fmt.Println("Inserido com sucesso")

	id, err := result.LastInsertId()
	if err != nil {
		return nil
	}
	return &models.Usuario{
		ID:    int(id),
		Nome:  nome,
		Email: email,
	}
}

func (s *UsuarioService) BuscarPorID(id int) *models.Usuario {
	var usuario models.Usuario
	var senhaHash string
	err := database.DB.QueryRow("SELECT id, nome, email, senha FROM usuarios WHERE id = ?", id).Scan(&usuario.ID, &usuario.Nome, &usuario.Email, &senhaHash)
	if err != nil {
		return nil
	}
	return &usuario
}

func (s *UsuarioService) ListarTodos() []models.Usuario {
	rows, err := database.DB.Query("SELECT id, nome, email, senha FROM usuarios")
	if err != nil {
		return []models.Usuario{}
	}
	defer rows.Close()

	var usuarios []models.Usuario
	for rows.Next() {
		var usuario models.Usuario
		var senhaHash string
		err := rows.Scan(&usuario.ID, &usuario.Nome, &usuario.Email, &senhaHash)
		if err != nil {
			continue
		}
		usuarios = append(usuarios, usuario)
	}

	return usuarios
}

func (s *UsuarioService) VerificarSenha(email, senha string) bool {
	var senhaHash string
	err := database.DB.QueryRow("SELECT senha FROM usuarios WHERE email = ?", email).Scan(&senhaHash)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(senha))
	return err == nil
}
