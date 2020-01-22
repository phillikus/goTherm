module goTherm/db

go 1.13

require (
	github.com/lib/pq v1.3.0
	goTherm/api/interfaces v0.0.0
	goTherm/config v0.0.0
)

replace goTherm/api/interfaces => ./../api/interfaces

replace goTherm/config => ./../config
