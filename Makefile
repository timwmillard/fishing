
all:

help: ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

test: ## Run the go test
	go test -cover ./...

testall: ## Run the integration, unit & race tests. Will create a docker continer for integration tests.
	go test -tags=integration -race -cover ./...

gen: ## Generate source code
	go generate ./...

build: ## Build the application 
	go build -o cmd/fishingd/fishingd ./cmd/fishingd

run: ## Run the application
	cmd/fishingd/fishingd

clean: ## Remove build binaries
	rm cmd/fishingd/fishingd

docker-db: ## Start the database using docker
	docker run --name fishing-db -p 5000:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=fish -e POSTGRES_DB=fishingcomp -d postgres:12-alpine

docker-psql: ## Connect to psql running in the docker container
	docker exec -it fishing-db psql -U root -d fishingcomp

psql-docker-db: ## Run the local psql connecting the docker database
	psql -h localhost -U root -p 5000 -d fishingcomp

docker-logs: ## Show the docker database logs
	docker logs fishing-db -f

docker-clean:  ## Remove the docker database and container
	-docker stop fishing-db
	-docker rm fishing-db

docker-migrate: docker-migrate-up ## Migrate up docker database

docker-migrate-up: ## Migrate up docker database
	migrate -path postgres/migrations -database "postgres://root:fish@localhost:5000/fishingcomp?sslmode=disable" up

docker-migrate-down: ## Migrate down docker database
	migrate -path postgres/migrations -database "postgres://root:fish@localhost:5000/fishingcomp?sslmode=disable" down

docker-db-reset: docker-clean docker-db sleep docker-migrate ## Remove the docker database and restart it

sleep:
	sleep 5
