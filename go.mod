module github.com/GodiusWick/GoLangZharkovValeriy

go 1.16

require github.com/lib/pq v1.10.2

require internal/Structures v1.0.0

replace internal/Structures => ./internal/Structures

require internal/Postgres v1.0.0

replace internal/Postgres => ./internal/Postgres

require internal/XMLReader v1.0.0

replace internal/XMLReader => ./internal/XMLReader
