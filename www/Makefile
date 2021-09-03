postgres:
	docker container run --name PosGo -P -e POSTGRES_USER=root -e POSTGRES_PASSWORD=12345 -d postgres:13.3

createdb:
	docker exec -it PostgreGo createdb --username=root --owner=root ProjectDB

dropdb:
	docker exec -it PostgreGo dropdb ProjectDB

migration:
	migrate -path db/migration -database "postgresql://root:12345@127.0.0.1:49161/ProjectDB?sslmode=disable" -verbose up

PHONY: createdb postgres dropdb migration
