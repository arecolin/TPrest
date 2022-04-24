module main

go 1.15

require github.com/gorilla/mux v1.8.0

require internal/entities v1.0.0

replace internal/entities => ./internal/entities

replace internal/bdd => ./internal/bdd

require (
	github.com/boltdb/bolt v1.3.1
	golang.org/x/sys v0.0.0-20220422013727-9388b58f7150 // indirect
	internal/bdd v0.0.0-00010101000000-000000000000
	internal/persistence v1.0.0
	internal/web/rest v0.0.0-00010101000000-000000000000
)

replace internal/persistence => ./internal/persistence

replace internal/web/rest => ./internal/web/rest
