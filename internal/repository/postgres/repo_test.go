// Package postgres provides
package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/tingchima/gogolook/configs"
)

var testDBConn *sqlx.DB
var testDBName string

const migrationPath = "file://../../../migrations"

func TestMain(m *testing.M) {

	cfg := configs.NewConfig("app-test")

	dbCfg := cfg.Database()

	testDBName = dbCfg.DBName

	conn, closeDB, err := setupTestDB(dbCfg)
	if err != nil {
		log.Printf("setup test db fail, err: " + err.Error())
		return
	}
	defer closeDB()

	testDBConn = conn
	_ = m.Run()
}

func getTestDBConn() *sqlx.DB {
	return testDBConn
}

func setupTestDB(cfg *configs.Database) (*sqlx.DB, func(), error) {

	var (
		dbConn  *sql.DB
		err     error
		closeDB func()
	)

	{
		connInfo := fmt.Sprintf(
			"user=%s password=%s host=%s sslmode=disable",
			cfg.Username,
			cfg.Password,
			cfg.Host,
		)

		dbConn, err = sql.Open("postgres", connInfo)
		if err != nil {
			return nil, nil, errors.WithMessage(err, "create test db connection fail")
		}

		createTestDBSql := fmt.Sprintf("CREATE DATABASE %s", cfg.DBName)

		_, err = dbConn.ExecContext(context.Background(), createTestDBSql)
		if err != nil {
			log.Println("exec create test db sql fail")
			return nil, nil, err
		}

		closeDB = func() {
			log.Println("start to close database")

			dropTestDBSql := fmt.Sprintf("DROP DATABASE %s WITH (FORCE)", cfg.DBName)

			_, err = dbConn.ExecContext(context.Background(), dropTestDBSql)
			if err != nil {
				log.Printf("exec drop test db sql fail, err: %s\n", err.Error())
				return
			}
		}
	}

	// create test db connection
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)

	testDBConn, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, nil, errors.WithMessage(err, "create test db connection fail")
	}

	// execute test db migration
	m, err := migrate.New(migrationPath, dsn)
	if err != nil {
		return nil, closeDB, errors.WithMessage(err, "migration process fail")
	}

	err = m.Up()
	if err != nil {
		return nil, closeDB, errors.WithMessage(err, "migrate test db fail")
	}

	return testDBConn, closeDB, nil
}

func setupTestData(sqlDB *sqlx.DB, files ...string) error {

	sql := `
		SELECT
			Concat('TRUNCATE TABLE ', TABLE_NAME)
		FROM
			INFORMATION_SCHEMA.TABLES
		WHERE
			table_schema = '%s';
	`

	sql = fmt.Sprintf(sql, testDBName)

	_, err := sqlDB.Exec(sql)
	if err != nil {
		log.Printf("fail to truncate all, err: %s", err.Error())
		return err
	}

	fixtures, err := testfixtures.New(
		testfixtures.Database(sqlDB.DB),    // You database connection
		testfixtures.Dialect("postgresql"), // Available: "postgresql", "timescaledb", "mysql", "mariadb", "sqlite" and "sqlserver"
		testfixtures.Files(files...),       // The directory containing the YAML files
		testfixtures.Location(time.UTC),
	)
	if err != nil {
		log.Printf("new test data loader fail, err: %s", err.Error())
		return err
	}

	err = fixtures.Load()
	if err != nil {
		log.Printf("load test data fail, err: %s", err.Error())
		return err
	}

	return nil
}
