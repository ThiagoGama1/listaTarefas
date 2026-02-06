package models

import "time"

type Tarefa struct {
	Id                int	`json:"id"`
	Nome              string	`json:"nome"`
	Custo             float64	`json:"custo"`
	DataLimite        time.Time	`json:"data_limite"`
	OrdemApresentacao int	`json:"ordem_apresentacao"`
}
