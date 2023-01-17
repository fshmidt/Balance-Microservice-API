build:
	docker-compose build balance-app

run:
	docker-compose up balance-app



migrates:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable' up
