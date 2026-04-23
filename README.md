# listaTarefas

Aplicação fullstack de gerenciamento de tarefas, desenvolvida com **Go (Golang)** no backend e **JavaScript puro** no frontend.

## 🚀 Tecnologias

- **Backend:** Go, Gin, pgx (PostgreSQL)
- **Banco de dados:** PostgreSQL
- **Frontend:** JavaScript, Bootstrap
- **Infraestrutura:** Docker, Docker Compose

## 📁 Estrutura do projeto

```
listaTarefas/
├── assets/js/     # Frontend (JavaScript puro)
├── cmd/           # Ponto de entrada da aplicação
├── data/          # Configuração e conexão com banco
├── handler/       # Handlers HTTP (rotas)
├── models/        # Estruturas de dados
├── templates/     # Templates HTML
├── Dockerfile
├── docker-compose.yml
└── go.mod
```

## ⚙️ Como rodar

Não é necessário ter Go ou PostgreSQL instalados. Basta ter o **Docker** e o **Docker Compose**.

```bash
# Clone o repositório
git clone https://github.com/ThiagoGama1/listaTarefas.git
cd listaTarefas

# Suba a aplicação
docker-compose up
```

A aplicação estará disponível em `http://localhost:8080`

## 💡 Decisões técnicas

- **Gin** foi escolhido pela leveza e familiaridade com a linguagem
- **pgx** para comunicação eficiente com o PostgreSQL
- **Docker** garante que a aplicação rode em qualquer máquina sem configuração manual
- Validações de segurança implementadas nos handlers para rejeitar dados inválidos
- Frontend em JS puro para manter o projeto leve, sem dependências de framework
