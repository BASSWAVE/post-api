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

postgres:
	sudo docker run \
	--name postgres-posts \
	-e POSTGRES_PASSWORD=pass \
	-e POSTGRES_USER=user \
	-p 5432:5432 \
	-d postgres