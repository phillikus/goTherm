module github.com/goTherm/api

go 1.13

require (
	github.com/goTherm/api/interfaces v0.0.0
	github.com/goTherm/config v0.0.0
	github.com/goTherm/db v0.0.0
)

replace github.com/goTherm/api/interfaces => ./interfaces

replace github.com/goTherm/db => ./../db

replace github.com/goTherm/config => ./../config
