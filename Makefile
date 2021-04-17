
all:

test:
	go test ./...

gen:
	go generate ./...

build:
	go build -o cmd/fishingd/fishingd ./cmd/fishingd

run:
	cmd/fishingd/fishingd

clean:
	rm cmd/fishingd/fishingd

docker-db:
	docker run --name fishing-db -p 5000:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=fish -e POSTGRES_DB=fishingcomp -d postgres:12-alpine

docker-psql:
	docker exec -it fishing-db psql -U root -d fishingcomp

psql-docker-db:
	psql -h localhost -U root -p 5000 -d fishingcomp

docker-logs:
	docker logs fishing-db -f

docker-clean:
	-docker stop fishing-db
	-docker rm fishing-db

docker-migrate: docker-migrate-up

docker-migrate-up:
	migrate -path postgres/migrations -database "postgres://root:fish@localhost:5000/fishingcomp?sslmode=disable" up

docker-migrate-down:
	migrate -path postgres/migrations -database "postgres://root:fish@localhost:5000/fishingcomp?sslmode=disable" down

docker-db-reset: docker-clean docker-db sleep docker-migrate

sleep:
	sleep 5