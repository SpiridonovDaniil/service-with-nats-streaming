package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"l0/internal/config"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Db struct {
	db *sqlx.DB
}

func New(cfg config.Postgres) *Db {
	conn, err := sqlx.Connect("postgres",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			cfg.User,
			cfg.Pass,
			cfg.Address,
			cfg.Port,
			cfg.Db,
		))
	if err != nil {
		log.Fatal(err)
	}

	return &Db{db: conn}
}

func (d *Db) InsertData(ctx context.Context, data json.RawMessage, id string) error {
	query := `INSERT INTO l0 (id, data) VALUES ($1, $2)`
	_, err := d.db.ExecContext(ctx, query, id, data)
	if err != nil {
		return fmt.Errorf("[insertData] %w", err)
	}

	return nil
}

func (d *Db) GetAll(ctx context.Context) (map[string]json.RawMessage, error) {
	resp := make(map[string]json.RawMessage)
	query := `SELECT * FROM l0`

	rows, err := d.db.QueryContext(ctx, query)
	if err != nil {
		return resp, fmt.Errorf("data could not be recovered, err: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var data json.RawMessage
		err := rows.Scan(&id, &data)
		if err != nil {
			return resp, fmt.Errorf("data could not be recovered, err: %w", err)
		}
		resp[id] = data
	}

	return resp, nil
}
