# Include variables from the .envrc file
include .envrc

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## db/migration/create name=$1: Creating migration file
.PHONY: db/migration/create
db/migration/create:
	@echo 'Creating migration files for ${name}...'
	migrate create -ext sql -dir ./migrations -seq ${name}

## db/migration/up: Running up migrations
.PHONY: db/migration/up
db/migration/up: confirm
	@echo 'Running up migrations...'
	migrate -path ./migrations -database ${DB_DSN} -verbose up

## db/migration/down: Running down migrations
.PHONY: db/migration/down
db/migration/down:
	@echo 'Running down migrations...'
	migrate -path ./migrations -database ${DB_DSN} -verbose down