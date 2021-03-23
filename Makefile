APP=HandsOnGoLang
APP_VERSION:=0.1
APP_EXECUTABLE="./out/$(APP)"

CONFIG_FILE="./.env"
HTTP_SERVE_COMMAND="http-serve"
GRPC_SERVE_COMMAND="grpc-serve"
MIGRATE_COMMAND="migrate"
ROLLBACK_COMMAND="rollback"

deps:
	go mod download
compile:
	mkdir -p out/
	go build -ldflags "-X main.version=$(APP_VERSION)" -o $(APP_EXECUTABLE) cmd/*.go

build: deps compile

httpserve: build
	$(APP_EXECUTABLE) -configFile=$(CONFIG_FILE) $(HTTP_SERVE_COMMAND)

init-db:
	psql -c "create user payment_user_go superuser password '${DB_PASSWORD}';" -U postgres
	psql -c "create database payment_db_go owner=payment_user_go" -U postgres

migrate: build
	$(APP_EXECUTABLE) -configFile=$(CONFIG_FILE) $(MIGRATE_COMMAND)