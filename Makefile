migrate_up:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up

migrate_down:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' down

run_db:
	docker run --name=coins-db -e POSTGRES_PASSWORD="qwerty" -p 5436:5432 -d --rm postgres

build:
	docker-compose build coins-app

run:
	docker-compose up coins-app