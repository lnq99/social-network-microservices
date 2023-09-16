DB_DRIVER=postgresql
USERNAME=program
PASSWORD=test
HOST=localhost
PORT=5432
DB_NAME=account

DB_URL="$(DB_DRIVER)://$(USERNAME):$(PASSWORD)@$(HOST):$(PORT)/$(DB_NAME)?sslmode=disable"

MIGRATION_DIR=db/migration

migrate_create:
#	migrate create -ext sql -dir $(MIGRATION_DIR) -seq init
#	migrate create -ext sql -dir $(MIGRATION_DIR) -seq seed
	migrate create -ext sql -dir $(MIGRATION_DIR) -seq func

migrate_up:
	migrate -path $(MIGRATION_DIR) -database $(DB_URL) -verbose up

migrate_up1:
	migrate -path $(MIGRATION_DIR) -database $(DB_URL) -verbose up 1

migrate_down:
	migrate -path $(MIGRATION_DIR) -database $(DB_URL) -verbose down

migrate_down1:
	migrate -path $(MIGRATION_DIR) -database $(DB_URL) -verbose down 1

sqlc:
	sqlc generate