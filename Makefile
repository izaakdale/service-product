postgres:
	docker run --name postDb -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=password -e POSTGRES_DB=ordering-app -d postgres
createdb:
	docker exec -it service-product_db_1 createdb --username=admin --owner=admin goBank
migrateup:
	migrate -path db/migration -database "postgresql://admin:password@localhost:5432/ordering-app?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://admin:password@localhost:5432/ordering-app?sslmode=disable" -verbose down
sql:
	sqlc generate
run:
	DB_HOST=localhost DB_PORT=5432 DB_USER=admin DB_PASSWORD=password DB_TABLE=ordering-app SERVICE_PORT=8080 go run main.go
build:
	GOOS=linux go build -o service-product . 
docker_build:
	docker build --no-cache -t service-product .