

build:
	go build .

run:
	go run .

test:
	go test .

lint:
	staticcheck ./...

migrate-up:
	migrate -path database/migrations -database "postgres://tim@localhost:5432/fishingcomp?sslmode=disable" up

migrate-down:
	migrate -path database/migrations -database "postgres://tim@localhost:5432/fishingcomp?sslmode=disable" down