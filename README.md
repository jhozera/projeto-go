# API de Usuários em Go

Projeto pequeno em Go: API para gerenciar usuários.

Passos mínimos para rodar:

1. Baixe dependências:

```bash
go mod download
```

2. Crie um arquivo `.env` na raiz com estas chaves:

```
DB_USER=
DB_PASSWORD=
DB_HOST=
DB_PORT=
DB_NAME=
```

3. Rode:

```bash
go run main.go
```

4. Teste com Postman ou curl:
- `GET /usuario`
- `GET /usuario?id=1`
- `POST /usuario` (body: `nome`, `email`, `senha`)
