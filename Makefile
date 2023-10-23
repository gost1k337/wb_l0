include .env
export

RUNNER=migrate

ifeq ($(p),host)
 	RUNNER=sql-migrate
endif

SOURCE="FILE://MIGRATIONS"
MIGRATE=$(RUNNER)
DB=${MIGRATION_DB_URL}

migrate-status:
	$(MIGRATE) version

migrate-up:
	$(MIGRATE) -source ${SOURCE} -database "${DB}" up

migrate-down:
	$(MIGRATE) -source ${SOURCE} -database "${DB}" down

docker-build:
	docker-compose -p us build

docker-run:
	docker-compose -p us up -d

docker-stop:
	docker-compose -p us stop

run-sub:
	cd subscriber && go run main.go

run-pub:
	cd publisher && go run main.go -c ${COUNT}
