package general

import (
	"context"
	"currency/internal/config"
	"database/sql"
	"embed"
	"fmt"
	"time"

	"github.com/pressly/goose/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Connection *sql.DB
	Gorm       *gorm.DB
)

//go:embed migrations/*.sql
var embedMigrations embed.FS //

func Connect() error {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Currency.MainDatabase.Username, config.Currency.MainDatabase.Password, config.Currency.MainDatabase.Host, config.Currency.MainDatabase.Port, config.Currency.MainDatabase.Base)

	Connection, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	Gorm, err = gorm.Open(mysql.New(mysql.Config{
		Conn: Connection,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return err
	}
	// Write function that sum

	Connection.SetConnMaxLifetime(0)
	Connection.SetMaxOpenConns(100)
	Connection.SetMaxIdleConns(100)

	goose.SetBaseFS(embedMigrations)
	goose.SetDialect("mysql")
	if err = goose.Up(Connection, "migrations"); err != nil {
		return err
	}

	return nil
}

func Exec(query string, args ...any) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := Connection.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("Exec: %s: error %s when preparing SQL statement", query, err)
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		return fmt.Errorf("Exec: %s error %s when inserting row into blacklist table", query, err)
	}
	return nil
}

func Close() {
	Connection.Close()
}
