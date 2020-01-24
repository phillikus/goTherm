module goTherm

go 1.13

require (
	github.com/goTherm/api v0.0.0
	github.com/goTherm/api/interfaces v0.0.0
	github.com/goTherm/config v0.0.0
	github.com/goTherm/db v0.0.0
)

replace github.com/goTherm/api/interfaces => ./api/interfaces

replace github.com/goTherm/api => ./api

replace github.com/goTherm/config => ./config

replace github.com/goTherm/db => ./db
