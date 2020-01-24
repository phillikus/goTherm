module goTherm

go 1.13

require (
	github.com/phillikus/goTherm/api v0.0.0
	github.com/phillikus/goTherm/api/interfaces v0.0.0
	github.com/phillikus/goTherm/config v0.0.0
	github.com/phillikus/goTherm/db v0.0.0
)

replace github.com/phillikus/goTherm/api/interfaces => ./api/interfaces

replace github.com/phillikus/goTherm/api => ./api

replace github.com/phillikus/goTherm/config => ./config

replace github.com/phillikus/goTherm/db => ./db
