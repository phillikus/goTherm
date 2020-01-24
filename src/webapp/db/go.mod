module github.com/phillikus/goTherm/db

go 1.13

require (
	github.com/lib/pq v1.3.0
	github.com/phillikus/goTherm/api/interfaces v0.0.0
	github.com/phillikus/goTherm/config v0.0.0
)

replace github.com/phillikus/goTherm/api/interfaces => ./../api/interfaces

replace github.com/phillikus/goTherm/config => ./../config
