
ENV ?= local


include .env



-include .env.$(ENV)

env:
	@echo ENV=$(ENV)
	@echo PGHOST=$(PGHOST)
	@echo PGPORT=$(PGPORT)
	@echo PGDATABASE=$(PGDATABASE)
	@echo PGUSER=$(PGUSER)
	@echo PGPASSWORD=$(PGPASSWORD)

all: build

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

run: build ## Run the application
	cmd/fishingd/fishingd

clean: ## Remove build binaries
	rm cmd/fishingd/fishingd

docker-db: ## Start the database using docker
	docker run --name $(PGDATABASE) -p $(PGPORT):5432 -e POSTGRES_USER=$(PGUSER) -e POSTGRES_PASSWORD=$(PGPASSWORD) -e POSTGRES_DB=$(PGDATABASE) -d postgres:13-alpine

docker-psql: ## Connect to psql running in the docker container
	docker exec -it $(PGDATABASE) psql -U $(PGUSER) -d $(PGDATABASE)

psql-docker-db: ## Run the local psql connecting the docker database
	psql -h localhost -U $(PGUSER) -p $(PGPORT) -d $(PGDATABASE)

docker-logs: ## Show the docker database logs
	docker logs $(PGDATABASE) -f

docker-clean:  ## Remove the docker database and container
	-docker stop $(PGDATABASE)
	-docker rm $(PGDATABASE)

docker-db-reset: docker-clean docker-db sleep migrate ## Remove the docker database and restart it

sleep:
	sleep 5

migrate: migrate-up ## Migrate up docker database

migrate-up: ## Migrate up docker database
	migrate -path postgres/migrations -database "postgres://$(PGUSER):$(PGPASSWORD)@$(PGHOST):$(PGPORT)/$(PGDATABASE)?sslmode=disable" up

migrate-down: ## Migrate down docker database
	migrate -path postgres/migrations -database "postgres://$(PGUSER):$(PGPASSWORD)@$(PGHOST):$(PGPORT)/$(PGDATABASE)?sslmode=disable" down
