module github.com/GodiusWick/GoLangZharkovValeriy

go 1.16

require github.com/lib/pq v1.10.2



require (
	postgres v0.0.0-00010101000000-000000000000
	structures v0.0.0-00010101000000-000000000000
	xmlreader v0.0.0-00010101000000-000000000000
)

replace xmlreader => ./xmlreader

replace postgres => ./postgres

replace structures => ./structures