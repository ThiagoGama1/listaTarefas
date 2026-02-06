package data

import (
	"context"
	"listaTarefas/models"

	"github.com/jackc/pgx/v5"
)

func ObterMaiorOrdem(conn *pgx.Conn) (int, error) {
	var maiorOrdem int

	query := `SELECT COALESCE(MAX(ordem_apresentacao), 0) FROM tarefas `

	err := conn.QueryRow(context.Background(), query).Scan(&maiorOrdem)
	if err != nil {
		return 0, err
	}
	return maiorOrdem, nil
}
func IncluirTarefa(conn *pgx.Conn, task models.Tarefa) error {
	var ultimoIndice int
	var err error
	ultimoIndice, err = ObterMaiorOrdem(conn)

	if err != nil {
		return err
	}
	var novaOrdem = ultimoIndice + 1

	query := `INSERT INTO tarefas(
	nome, custo, data_limite, ordem_apresentacao)
	VALUES($1, $2, $3, $4)`

	_, err = conn.Exec(context.Background(), query,
		task.Nome,
		task.Custo,
		task.DataLimite,
		novaOrdem,
	)
	return err
}

func ListarTarefas(conn *pgx.Conn) ([]models.Tarefa, error) {
	query := `SELECT * FROM tarefas
			  ORDER BY ordem_apresentacao`

	rows, err := conn.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	listaCompleta := []models.Tarefa{}
	for rows.Next() {
		var tarefa models.Tarefa
		err = rows.Scan(&tarefa.Id, &tarefa.Nome, &tarefa.Custo, &tarefa.DataLimite, &tarefa.OrdemApresentacao)
		if err != nil {
			return nil, err
		}
		listaCompleta = append(listaCompleta, tarefa)
	}

	return listaCompleta, nil
}

func AtualizarTarefa(conn *pgx.Conn, task models.Tarefa) error {
	query := `UPDATE tarefas
SET nome = $1,
custo = $2,
data_limite = $3
WHERE id = $4`

	_, err := conn.Exec(context.Background(), query,
		task.Nome,
		task.Custo,
		task.DataLimite,
		task.Id,
	)

	return err

}
func ExcluirTarefa(conn *pgx.Conn, id int) error {
	query := `DELETE FROM tarefas WHERE id = $1`

	_, err := conn.Exec(context.Background(), query, id)

	return err
}
func VerificarNomeExiste(conn *pgx.Conn, nome string) bool {
	var quantidade int
	query := `SELECT COUNT(nome) FROM tarefas WHERE nome = $1`
	err := conn.QueryRow(context.Background(), query, nome).Scan(&quantidade)
	if err != nil {
		return false
	}
	if quantidade != 0 {
		return true
	}

	return false
}

func VerificarNomeEdicao(conn *pgx.Conn, nome string, id int) bool{
	var quantidade int
	query := `SELECT COUNT(nome) FROM tarefas WHERE nome = $1 AND id != $2`
	err := conn.QueryRow(context.Background(), query, nome, id).Scan(&quantidade)
	if err != nil {
		return false
	}
	if quantidade != 0 {
		return true
	}

	return false
}

func SubirTarefa(conn *pgx.Conn, id int) (bool, error) {
	var ordemAtual int
	var ordemDeCima int
	var idDeCima int
	var ordemTemporaria = -1
	query := `SELECT ordem_apresentacao FROM tarefas WHERE id = $1`

	err := conn.QueryRow(context.Background(), query, id).Scan(&ordemAtual)
	if err != nil{
		return false, err 
	}

	
	query2 := `SELECT ordem_apresentacao, id FROM tarefas WHERE ordem_apresentacao < $1
	ORDER BY ordem_apresentacao DESC LIMIT 1`
	err = conn.QueryRow(context.Background(), query2, ordemAtual).Scan(&ordemDeCima, &idDeCima)
	if err == pgx.ErrNoRows {
		return true, nil
	}

	queryPosicaoTemp := `UPDATE tarefas
	SET ordem_apresentacao = $1
	WHERE id = $2`

	_, err = conn.Exec(context.Background(), queryPosicaoTemp, ordemTemporaria, idDeCima)

	if err != nil {
		return false, err
	}
	
	queryDeslocamento := `UPDATE tarefas
	SET ordem_apresentacao = $1
	WHERE id = $2`

	_, err = conn.Exec(context.Background(), queryDeslocamento, ordemDeCima, id)

	if err != nil{
		return false, err
	}

	queryRealocacao := `UPDATE tarefas
	SET ordem_apresentacao = $1
	WHERE id = $2`

	_, err = conn.Exec(context.Background(), queryRealocacao, ordemAtual, idDeCima)

	if err != nil{
		return false, err
	}
	
	return true, nil

}

func DescerTarefa(conn *pgx.Conn, id int) (bool, error){
	var idDeBaixo int
	var ordemAtual int
	var ordemDeBaixo int
	var ordemTemporaria = -1

	queryOrdemAtual := `SELECT ordem_apresentacao FROM tarefas
						WHERE id = $1`

	err := conn.QueryRow(context.Background(), queryOrdemAtual, id).Scan(&ordemAtual)

	if err != nil{
		return false, err
	}

	queryOrdemEIdDeBaixo := `SELECT ordem_apresentacao, id
	FROM tarefas WHERE ordem_apresentacao > $1
	ORDER BY ordem_apresentacao ASC LIMIT 1`

	err = conn.QueryRow(context.Background(), queryOrdemEIdDeBaixo, ordemAtual).Scan(&ordemDeBaixo, &idDeBaixo)

	if err == pgx.ErrNoRows{
		return true, nil
	}

	queryUpdateTemporario := `UPDATE tarefas
	SET ordem_apresentacao = $1
	WHERE id = $2`

	_, err = conn.Exec(context.Background(), queryUpdateTemporario, ordemTemporaria, idDeBaixo)
	
	if err != nil {
		return false, err
	}


	queryDeslocamento := `UPDATE tarefas
	SET ordem_apresentacao = $1
	WHERE id = $2`

	_, err = conn.Exec(context.Background(), queryDeslocamento, ordemDeBaixo, id)

	if err != nil {
		return false, err
	}

	queryRealocacao := `UPDATE tarefas
	SET ordem_apresentacao = $1
	WHERE id = $2`

	_, err = conn.Exec(context.Background(), queryRealocacao, ordemAtual, idDeBaixo)

	if err != nil {
		return false, err
	}

	return true, nil

}
