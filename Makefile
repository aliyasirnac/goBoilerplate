run:
	@go run cmd/api/main.go

up:
	@docker-compose up -d

down:
	@docker-compose down

