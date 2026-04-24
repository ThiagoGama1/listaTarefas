# listaTarefas

Aplicação fullstack de gerenciamento de tarefas com API REST em Go, banco de dados PostgreSQL e containerização completa via Docker. O projeto permite criar, visualizar, editar, reordenar e excluir tarefas, com validações no backend para garantir a integridade dos dados.

---

## Stack técnica

| Camada | Tecnologia |
|---|---|
| Backend | Go + Gin |
| Banco de dados | PostgreSQL (driver pgx) |
| Frontend | JavaScript puro + Bootstrap |
| Containerização | Docker + Docker Compose |

---

## Arquitetura

O projeto segue uma estrutura em camadas com responsabilidades bem separadas:

```
listaTarefas/
├── cmd/            # Ponto de entrada da aplicação (main.go)
├── data/           # Conexão e queries ao banco de dados (PostgreSQL via pgx)
├── handler/        # Handlers HTTP das rotas (lógica de negócio e validações)
├── models/         # Structs de domínio (Tarefa: id, nome, custo, data_limite, ordem)
├── templates/      # Templates HTML servidos pelo backend
├── assets/js/      # JavaScript do frontend (requisições à API, manipulação do DOM)
├── Dockerfile
├── docker-compose.yml
└── go.mod
```

**Modelo de dados — Tarefa:**
```go
type Tarefa struct {
    Id                int       `json:"id"`
    Nome              string    `json:"nome"`
    Custo             float64   `json:"custo"`
    DataLimite        time.Time `json:"data_limite"`
    OrdemApresentacao int       `json:"ordem_apresentacao"`
}
```

---

## Funcionalidades

- Listagem de tarefas ordenadas por `ordem_apresentacao`
- - Cadastro de nova t
