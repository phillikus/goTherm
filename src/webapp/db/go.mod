module goTherm/db

go 1.13

require (
	github.com/lib/pq v1.3.0
	github.com/phillikus/goTherm v0.0.0-20191211184658-cfca36d382b0
	goTherm/api/interfaces v0.0.0
)

replace goTherm/api/interfaces => ./../api/interfaces
