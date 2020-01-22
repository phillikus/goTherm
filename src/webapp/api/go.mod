module goTherm/api

go 1.13

require (
	goTherm/api/interfaces v0.0.0
	goTherm/config v0.0.0
	goTherm/db v0.0.0
)

replace goTherm/api/interfaces => ./interfaces

replace goTherm/db => ./../db

replace goTherm/config => ./../config
