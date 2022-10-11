createdb:
	docker exec -it service-product_db_1 createdb --username=admin --owner=admin goBank
migrateup:
	migrate -path db/migration -database "postgresql://admin:password@localhost:5432/ordering-app?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://admin:password@localhost:5432/ordering-app?sslmode=disable" -verbose down
sql:
	sqlc generate
run:
	DB_HOST=localhost DB_PORT=5432 DB_USER=admin DB_PASSWORD=password DB_TABLE=ordering-app go run main.go