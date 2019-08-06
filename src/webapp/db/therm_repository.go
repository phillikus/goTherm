package db

import (
	"github.com/phillikus/goTherm/src/webapp/api/interfaces"
	"github.com/phillikus/goTherm/src/webapp/config"
)

type ThermRepository struct {
	config *config.Config
}

func (repository *ThermRepository) Insert(data api_interfaces.ThermData) {
	return
}

func (repository *ThermRepository) GetLatest(count int) []api_interfaces.ThermData {
	return make([]api_interfaces.ThermData, count)
}
