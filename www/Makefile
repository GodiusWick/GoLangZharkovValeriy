postgres:
	docker run --name PostgreGo -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=10IMclass -d postgres:13.3

createdb:
	docker exec -it PostgreGo createdb --username=root --owner=root TestGoDocker

dropdb:
	docker exec -it PostgreGo dropdb TestGoDocker

PHONY: createdb
