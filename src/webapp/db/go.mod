module github.com/goTherm/db

go 1.13

require (
	github.com/goTherm/api/interfaces v0.0.0
	github.com/goTherm/config v0.0.0
	github.com/lib/pq v1.3.0
)

replace github.com/goTherm/api/interfaces => ./../api/interfaces

replace github.com/goTherm/config => ./../config
