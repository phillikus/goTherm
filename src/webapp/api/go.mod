module github.com/phillikus/goTherm/api

go 1.13

require (
	github.com/phillikus/goTherm/api/interfaces v0.0.0
	github.com/phillikus/goTherm/config v0.0.0
	github.com/phillikus/goTherm/db v0.0.0
)

replace github.com/phillikus/goTherm/api/interfaces => ./interfaces

replace github.com/phillikus/goTherm/db => ./../db

replace github.com/phillikus/goTherm/config => ./../config
