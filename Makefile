ifeq ($(POSTGRES_SETUP_TEST),)
	POSTGRES_SETUP_TEST := user=user password=pass dbname=postgres host=localhost port=5432 sslmode=disable
endif

MIGRATION_FOLDER=$(CURDIR)/migrations

migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" up

migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" down

docker-up:
	 docker compose up -d

docker-down:
	docker compose down
