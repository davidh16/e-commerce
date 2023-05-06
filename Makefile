# Load environment variables from .env file
include .env
export $(shell sed 's/=.*//' .env)

.PHONY: create-database
.PHONY: drop-database
.PHONY: create-migration

create-database:
	@if psql -lqt | cut -d \| -f 1 | grep -qw $(DB_NAME); then \
		echo "Database already exists"; \
	else \
		createdb -h $(DB_HOST) -p $(shell expr $(DB_PORT) + 0) -U $(DB_USER) $(DB_NAME); \
	echo "Database created"; \
	fi

drop-database:
	@if ! psql -lqt | cut -d \| -f 1 | grep -qw $(DB_NAME); then \
		echo "Database does not exist"; \
	else \
		dropdb -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) $(DB_NAME); \
		echo "Database dropped"; \
	fi

create-migration:
	$(eval NAME := $(filter-out $@,$(MAKECMDGOALS)))
	$(eval FILENAME := $(shell date +'%Y%m%d%H%M%S')_$(name))
	./migration_template.sh $(FILENAME) $(name)

migrate:
	migrate -path ./migrations -database "postgres://$(DB_USER)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up

rollback:
	migrate -path ./migrations -database "postgres://$(DB_USER)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" down -all