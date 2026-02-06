package data

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConectarBanco() (*pgx.Conn, error) {

	url := os.Getenv("DATABASE_URL")

	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// ... func CriaTabela continua igual ...
func CriaTabela(conn *pgx.Conn) error {
	query := `CREATE TABLE IF NOT EXISTS tarefas(
		id SERIAL PRIMARY KEY,
		nome VARCHAR(255) NOT NULL UNIQUE,
		custo DECIMAL(10, 2) NOT NULL CHECK (custo >= 0),
        data_limite DATE NOT NULL,
        ordem_apresentacao INTEGER NOT NULL UNIQUE
	)`
	_, err := conn.Exec(context.Background(), query)

	if err != nil {
		return err
	}
	return nil
}
