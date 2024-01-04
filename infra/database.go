// Package infra provides
package infra

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/tingchima/gogolook/configs"
)

// MustNewPostgresConn .
func MustNewPostgresConn(cfg *configs.Database) *sqlx.DB {

	conn := sqlx.MustOpen("postgres", resolvePostgresDSN(cfg))

	err := conn.Ping()
	if err != nil {
		log.Panicf("new pg fail, err: %s\n", err.Error())
	}

	log.Println("new pg successfully")
	return conn
}

// NewPostgresConn .
func NewPostgresConn(cfg *configs.Database) (*sqlx.DB, error) {

	conn, err := sqlx.Open("postgres", resolvePostgresDSN(cfg))
	if err != nil {
		log.Printf("new postgres connection fail, err: %s", err.Error())
		return nil, err
	}

	return conn, nil
}

// resolvePostgresDSN .
func resolvePostgresDSN(cfg *configs.Database) string {

	log.Println(fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		// "postgresql://%s:%s/%s?sslmode=disable",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	))

	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		// "postgresql://%s:%s/%s?sslmode=disable",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)
}
