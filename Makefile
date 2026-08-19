run:
	go run ./cmd/api
migrate-create:
	migrate create -ext sql -dir migrations -seq ${name}
migrate-up:
	migrate -path migrations -database "postgres://postgres:postgres@localhost:5432?restaurant_management?sslmode=disable" up
