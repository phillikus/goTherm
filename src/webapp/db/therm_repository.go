package db

import (
	"database/sql"
	api "github.com/phillikus/goTherm/api/interfaces"
	config "github.com/phillikus/goTherm/config"
	"time"

	_ "github.com/lib/pq"
)

type ThermRepository struct {
	Config *config.Config
}

type Temperature struct {
}

func (repository *ThermRepository) Insert(data api.ThermData) {
	data.Id = 0

	db := getDbConnection(repository.Config.Database.ConnectionString)
	defer db.Close()

	sql := "INSERT INTO dbo.temperature(temperature, create_date) VALUES($1, $2)"
	_, err := db.Exec(sql, data.Temperature, data.CreateDate)

	if err != nil {
		panic(err)
	}
}

func (repository *ThermRepository) GetLatest(count int) []api.ThermData {
	var (
		id          int
		temperature float32
		create_date time.Time
	)

	db := getDbConnection(repository.Config.Database.ConnectionString)
	defer db.Close()

	sql := "SELECT * FROM dbo.temperature ORDER BY create_date DESC LIMIT $1"
	rows, err := db.Query(sql, count)

	if err != nil {
		panic(err)
	}

	var entries []api.ThermData

	for rows.Next() {
		err := rows.Scan(&id, &temperature, &create_date)
		if err != nil {
			panic(err)
		}

		entries = append(entries, api.ThermData{
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
