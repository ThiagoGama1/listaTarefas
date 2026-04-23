# listaTarefas

Projeto fullstack de lista de tarefas feito em Go. Desenvolvi pra praticar a construção de uma API do zero com banco de dados real e subir tudo via Docker.

## Stack

- Go + Gin no backend
- PostgreSQL como banco (usando pgx)
- JavaScript puro e Bootstrap no frontend
- Docker e Docker Compose pra rodar tudo junto

## Estrutura

```
listaTarefas/
├── assets/js/     # frontend
├── cmd/           # ponto de entrada
├── data/          # conexão com o banco
├── handler/       # handlers das rotas
├── models/        # structs
├── templates/     # HTML
├── Dockerfile
├── docker-compose.yml
└── go.mod
```

## Como rodar

Precisa ter Docker instalado. Não precisa instalar Go nem PostgreSQL na máquina.

```bash
git clone https://github.com/ThiagoGama1/listaTarefas.git
cd listaTarefas
docker-compose up
```

Acessa em `http://localhost:8080`

## Por que essas escolhas

Usei Gin por ser leve e já ter alguma familiaridade. O pgx lida bem com PostgreSQL em Go. Coloquei validações nos handlers pra não deixar dados ruins entrarem no banco. Frontend simples de propósito — não queria adicionar complexidade onde não era necessário.
