module goTherm

go 1.13

require (
	goTherm/api v0.0.0
	goTherm/api/interfaces v0.0.0
	goTherm/config v0.0.0
	goTherm/db v0.0.0
)

replace goTherm/api/interfaces => ./api/interfaces

replace goTherm/api => ./api

replace goTherm/config => ./config

replace goTherm/db => ./db
