createdb:
	docker exec -it service-product_db_1 createdb --username=admin --owner=admin goBank
migrateup:
	migrate -path db/migration -database "postgresql://admin:password@localhost:5432/ordering-app?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://admin:password@localhost:5432/ordering-app?sslmode=disable" -verbose down
sql:
	sqlc generate