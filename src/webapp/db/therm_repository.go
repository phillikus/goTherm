package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/phillikus/goTherm/src/webapp/api/interfaces"
	"github.com/phillikus/goTherm/src/webapp/config"
	"time"
)

type ThermRepository struct {
	Config *config.Config
}

type Temperature struct {
}

func (repository *ThermRepository) Insert(data api_interfaces.ThermData) {
	data.Id = 0

	db := getDbConnection(repository.Config.Database.ConnectionString)
	defer db.Close()

	sql := "INSERT INTO dbo.temperature(temperature, create_date) VALUES($1, $2)"
	_, err := db.Exec(sql, data.Temperature, data.CreateDate)

	if err != nil {
		panic(err)
	}
}

func (repository *ThermRepository) GetLatest(count int) []api_interfaces.ThermData {
	var (
		id          int
		temperature float32
		create_date time.Time
	)

	var entries []api_interfaces.ThermData

	db := getDbConnection(repository.Config.Database.ConnectionString)
	defer db.Close()

	sql := "SELECT * FROM dbo.temperature ORDER BY create_date DESC LIMIT $1"
	rows, err := db.Query(sql, count)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err := rows.Scan(&id, &temperature, &create_date)
		if err != nil {
			panic(err)
		}
		entries = append(entries, api_interfaces.ThermData{
			Id:          id,
			Temperature: temperature,
			CreateDate:  create_date,
		})
	}

	return entries
}

func getDbConnection(connectionString string) *sql.DB {
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}

	return db
}
