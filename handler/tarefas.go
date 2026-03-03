package handler

import (
	"listaTarefas/data"
	"listaTarefas/models"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type TarefasHandler struct {
	Db *pgx.Conn
}

type TarefaInput struct {
	Nome       string      `json:"nome"`
	Custo      interface{} `json:"custo"`
	DataLimite time.Time   `json:"data_limite"`
}

func converterCusto(v interface{}) float64 {
	switch val := v.(type) {
	case float64:
		return val
	case string:
		val = strings.ReplaceAll(val, ",", ".")
		c, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return 0
		}
		return c
	default:
		return 0
	}
}

func (h *TarefasHandler) HandlerListarTarefas(c *gin.Context) {
	lista, err := data.ListarTarefas(h.Db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "falha ao carregar banco de dados"})
		return
	}

	c.JSON(http.StatusOK, lista)

}

func (h *TarefasHandler) HandlerIncluirTarefa(c *gin.Context) {
	var input TarefaInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "falha ao ler json"})
		return
	}

	var novaTarefa models.Tarefa
	var valorLimite = 999999999
	novaTarefa.Nome = input.Nome
	novaTarefa.Custo = converterCusto(input.Custo)
	novaTarefa.DataLimite = input.DataLimite

	var erros []string

	if novaTarefa.Nome == "" {
		erros = append(erros, "O nome da tarefa é obrigatório")

	} else if len(novaTarefa.Nome) > 50 {
		erros = append(erros, "O nome da tarefa deve ter no máximo 50 caracteres")

	} else if data.VerificarNomeExiste(h.Db, novaTarefa.Nome) {
		erros = append(erros, "O nome da tarefa já existe")
	}
	if novaTarefa.Custo < 0 {
		erros = append(erros, "O custo não pode ser negativo")

	} else if novaTarefa.Custo > float64(valorLimite) {
		erros = append(erros, "O custo da tarefa é muito alto")
	}
	if novaTarefa.DataLimite.IsZero() {
		erros = append(erros, "A data limite é obrigatória")
	}

	if len(erros) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"erro": strings.Join(erros, " | ")})
		return
	}

	err := data.IncluirTarefa(h.Db, novaTarefa)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "falha ao incluir tarefa"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"mensagem": "incluído com sucesso"})

}

func (h *TarefasHandler) HandlerEditarTarefa(c *gin.Context) {
	idTarefa := c.Param("id")
	id, err := strconv.Atoi(idTarefa)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "falha ao converter id"})
		return
	}

	var input TarefaInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "falha ao ler dados da tarefa"})
		return
	}

	var tarefaEditada models.Tarefa
	var valorLimite = 999999999
	tarefaEditada.Nome = input.Nome
	tarefaEditada.Custo = converterCusto(input.Custo)
	tarefaEditada.DataLimite = input.DataLimite

	var erros []string

	if tarefaEditada.Nome == "" {
		erros = append(erros, "O nome da tarefa é obrigatório")

	} else if len(tarefaEditada.Nome) > 50 {
		erros = append(erros, "O nome da tarefa deve ter no máximo 50 caracteres")

	} else if data.VerificarNomeEdicao(h.Db, tarefaEditada.Nome, tarefaEditada.Id) {
		erros = append(erros, "O nome da tarefa já existe")
	}
	if tarefaEditada.Custo < 0 {
		erros = append(erros, "O custo não pode ser negativo")

	} else if tarefaEditada.Custo > float64(valorLimite) {
		erros = append(erros, "O custo da tarefa é muito alto")
	}
	if tarefaEditada.DataLimite.IsZero() {
		erros = append(erros, "A data limite é obrigatória")
	}

	if len(erros) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"erro": strings.Join(erros, " | ")})
		return
	}
	tarefaEditada.Id = id

	err = data.AtualizarTarefa(h.Db, tarefaEditada)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "falha ao atualizar tarefa"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "tarefa editada com sucesso"})
}

func (h *TarefasHandler) HandlerExcluirTarefa(c *gin.Context) {
	idTarefa := c.Param("id")
	id, err := strconv.Atoi(idTarefa)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "falha ao converter id"})
		return
	}

	err = data.ExcluirTarefa(h.Db, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "falha ao deletar tarefa"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensagem": "tarefa excluída"})
}

func (h *TarefasHandler) HandlerSubirTarefa(c *gin.Context) {
	idTarefa := c.Param("id")
	id, err := strconv.Atoi(idTarefa)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "falha ao converter id"})
		return
	}
	var jaNoTopo bool
	jaNoTopo, err = data.SubirTarefa(h.Db, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "falha ao subir tarefa"})
		return
	}
	if jaNoTopo {
		c.JSON(http.StatusOK, gin.H{"mensagem": "tarefa já está no topo"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensagem": "tarefa elevada"})
}

func (h *TarefasHandler) HandlerDescerTarefa(c *gin.Context) {
	idTarefa := c.Param("id")
	id, err := strconv.Atoi(idTarefa)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "falha ao converter id"})
		return
	}
	_, err = data.DescerTarefa(h.Db, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "falha ao descer tarefa"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensagem": "tarefa rebaixada"})
}
