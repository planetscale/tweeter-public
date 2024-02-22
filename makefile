-include .env
export

setup: 
	bin/setup

run:
	bin/rails server

seed:
	cd faker;go mod tidy;go run main.go

migrate:
	rails psdb:migrate