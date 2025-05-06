package repository

import (
	"currency/app/repository/bigdata"
	"currency/app/repository/general"
	"currency/internal/logs"
)

func Connect() {
	err := general.Connect()
	if err != nil {
		logs.Fatal("Cant connect to main (mysql) repository", err)
	}

	err = bigdata.Connect()
	if err != nil {
		logs.Fatal("Cant connect to bigdata (clickhouse) repository", err)
	}
	logs.Info("Connected to main (mysql) and bigdata (clickhouse) repositories")
}

func Close() {
	general.Close()
	bigdata.Close()
}
