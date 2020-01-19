package api_interfaces

import "time"

type ThermData struct {
	Id          int
	Temperature float32
	CreateDate  time.Time
}

type ThermRepository interface {
	Insert(data ThermData)
	GetLatest(count int) []ThermData
}
