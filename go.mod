module github.com/GodiusWick/GoLangZharkovValeriy

go 1.16

require github.com/lib/pq v1.10.2
require internal/DataStructures v1.0.0
replace internal/DataStructures => ./internal/DataStructures
require internal/PostgresPars v1.0.0
replace internal/PostgresPars => ./internal/PostgresPars
require internal/XMLReader v1.0.0
replace internal/XMLReader => ./internal/XMLReader