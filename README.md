# Sistema de Lista de Tarefas

Aplicação web desenvolvida para cadastro, listagem, edição, exclusão e reordenação de tarefas, utilizando persistência em banco de dados PostgreSQL.

O projeto está disponível em ambiente de produção para testes imediatos, conforme solicitado.

**Link de Acesso:** https://listatarefas-production-12df.up.railway.app/

---

## Stack Tecnológica

* **Linguagem:** Go (Golang) 1.24
* **Framework Web:** Gin
* **Banco de Dados:** PostgreSQL (Driver pgx)
* **Frontend:** HTML5, JavaScript (Vanilla) e Bootstrap 5
* **Infraestrutura:** Docker e Docker Compose

## Decisões de Arquitetura

O desenvolvimento seguiu princípios de simplicidade e performance, focando estritamente nos requisitos funcionais.

1.  **Backend em Go:** A escolha da linguagem Go junto ao framework Gin visou otimizar o desempenho do sistema com baixo consumo de recursos. O código foi estruturado em pacotes distintos (handlers, models e database) para facilitar a manutenção e legibilidade.

2.  **Frontend Leve:** A implementação utilizou apenas JavaScript puro e Bootstrap, evitando a complexidade desnecessária de frameworks SPA. Isso garantiu a implementação de todas as regras de negócio (ordenação, validações e cálculo de custos) com um tempo de carregamento reduzido.

3.  **Dockerização:** O ambiente foi configurado via Docker para garantir consistência entre desenvolvimento e produção. O Dockerfile utiliza *multi-stage builds* para gerar uma imagem final eficiente e leve.

## Funcionalidades Atendidas

A aplicação cumpre integralmente os requisitos do desafio:

* Interface responsiva para listagem e gerenciamento de tarefas.
* Validação de campos obrigatórios e unicidade de nomes.
* Destaque visual automático para tarefas com custo superior a R$ 1.000,00.
* Sistema de reordenação de prioridade das tarefas.

## Execução Local

Para rodar o projeto localmente:

1.  Clone o repositório.
2.  Execute o comando:
    ```bash
    docker compose up --build
    ```
3.  Acesse `http://localhost:8080`.
