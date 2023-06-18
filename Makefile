# Load environment variables from .env file
include .env
export $(shell sed 's/=.*//' .env)

.PHONY: create-database
.PHONY: drop-database
.PHONY: create-migration

create-database:
	@if PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -lqt | cut -d \| -f 1 | grep -qw $(DB_NAME); then \
		echo "Database already exists"; \
	else \
	  	PGPASSWORD=$(DB_PASSWORD) createdb -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) $(DB_NAME); \
		echo "Database created"; \
	fi

drop-database:
	@if ! PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -lqt | cut -d \| -f 1 | grep -qw $(DB_NAME); then \
		echo "Database does not exist"; \
	else \
		PGPASSWORD=$(DB_PASSWORD) dropdb -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) $(DB_NAME); \
		echo "Database dropped"; \
	fi

create-migration:
	$(eval NAME := $(filter-out $@,$(MAKECMDGOALS)))
	$(eval FILENAME := $(shell date +'%Y%m%d%H%M%S')_$(name))
	./dev-tools/migration_template.sh $(FILENAME) $(name)

migrate:
	PGPASSWORD=$(DB_PASSWORD) migrate -path ./migrations -database "postgres://$(DB_USER)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up

rollback:
	PGPASSWORD=$(DB_PASSWORD) migrate -path ./migrations -database "postgres://$(DB_USER)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" down -all

up:
	docker-compose -f docker-compose-db.yml up --build -d
	docker-compose -f docker-compose.yml  up --build

down:
	docker-compose -f docker-compose-db.yml stop
	docker-compose -f docker-compose.yml down
	docker-compose -f docker-compose-db.yml down

stop:
	docker stop e-commerce

start:
	docker start e-commerce