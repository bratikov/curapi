package bigdata

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"reflect"

	"currency/internal/config"
	"time"

	"sync"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/elliotchance/orderedmap/v2"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS //

var (
	Connection driver.Conn
	mu         = &sync.Mutex{}
)

const debug = false

func Connect() error {
	var err error
	Connection, err = clickhouse.Open(&clickhouse.Options{
		Addr: []string{config.Currency.LogsDatabase.Host + ":" + config.Currency.LogsDatabase.Port},
		Auth: clickhouse.Auth{
			Database: config.Currency.LogsDatabase.Base,
			Username: config.Currency.LogsDatabase.Username,
			Password: config.Currency.LogsDatabase.Password,
		},
		Debug: debug,
		Debugf: func(format string, v ...any) {
			fmt.Printf(format, v)
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		MaxOpenConns:     1000,
		MaxIdleConns:     300,
		ConnMaxLifetime:  time.Duration(20) * time.Minute,
		ConnOpenStrategy: clickhouse.ConnOpenInOrder,
		//  BlockBufferSize:      10,
		//  MaxCompressionBuffer: 10240,
		ClientInfo: clickhouse.ClientInfo{ // optional, please see Client info section in the README.md
			Products: []struct {
				Name    string
				Version string
			}{
				{Name: "Currency-client", Version: "1.0"},
			},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to connect clickhouse server, err=%s", err)
	}

	if err = Connection.Ping(context.Background()); err != nil {
		return fmt.Errorf("failed to ping clickhouse server, err=%s", err)

	}

	if err := migrate(); err != nil {
		return fmt.Errorf("failed to migrate clickhouse server, err=%s", err)
	}

	return nil
}

func migrate() error {
	clickhouseURL := "clickhouse://" + config.Currency.LogsDatabase.Username + ":" + config.Currency.LogsDatabase.Password + "@" + config.Currency.LogsDatabase.Host + ":" + config.Currency.LogsDatabase.Port + "/" + config.Currency.LogsDatabase.Base

	db, err := sql.Open("clickhouse", clickhouseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	goose.SetBaseFS(embedMigrations)
	goose.SetDialect("clickhouse")
	if err = goose.Up(db, "migrations"); err != nil {
		return err
	}
	return nil
}

func GetRows(list driver.Rows) ([]map[string]interface{}, error) {
	types := list.ColumnTypes()
	var rows []map[string]interface{}
	for list.Next() {
		values := make([]interface{}, len(types))
		for i, t := range types {
			mu.Lock()
			values[i] = reflect.New(t.ScanType()).Interface()
			mu.Unlock()
		}

		if err := list.Scan(values...); err != nil {
			return nil, err
		}

		row := make(map[string]interface{})
		for i, t := range types {
			row[t.Name()] = values[i]
		}
		rows = append(rows, row)
	}
	return rows, nil
}

func GetRowsOrdered(list driver.Rows) ([]*orderedmap.OrderedMap[string, interface{}], error) {
	types := list.ColumnTypes()
	var rows []*orderedmap.OrderedMap[string, interface{}]
	for list.Next() {
		values := make([]interface{}, len(types))
		for i, t := range types {
			values[i] = reflect.New(t.ScanType()).Interface()
		}

		if err := list.Scan(values...); err != nil {
			return nil, err
		}

		m := orderedmap.NewOrderedMap[string, interface{}]()
		for i, t := range types {
			m.Set(t.Name(), values[i])
		}
		rows = append(rows, m)
	}
	return rows, nil
}

func Close() {
	Connection.Close()
}
