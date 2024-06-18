docker-up:
	docker compose -f deployments/docker-compose.yml up -d
docker-down:
	docker compose -f deployments/docker-compose.yml down 
docker-logs:
	docker compose -f deployments/docker-compose.yml logs -f
postgres-connect:
	docker exec -it my-postgres psql -U user mydb
run:
	go run cmd/go-back-end/main.go
example-user:
	curl -X POST http://localhost:8080/users \
                            -H 'Content-Type: application/json' \
                            -d '{"username": "lallalala", "email": "djafklsjklf@dfjkhlfdjkl.com"}'



