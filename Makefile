ENTRY_POINT=cmd/app
MIGRATE_POINT=cmd/migrate
MIGRATIONS_DIR=migrations

# Генерация контрактов
proto-generate:
	protoc --proto_path=proto \
		--go_out=gen/auth --go_opt=paths=source_relative \
		--go-grpc_out=gen/auth --go-grpc_opt=paths=source_relative \
		proto/auth.proto

# Генерация секретного ключа (пока тестовое исполнение)
generate-key:
	go run cmd/generate-key/generate_key.go

# Запуск
run:
	go run $(ENTRY_POINT)/main.go

# Запуск в тестовом режиме
run-dev:
	go run $(ENTRY_POINT)/main.go --dev

# Создание миграций
migrations-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $$name

# Применение миграций
migrate-up:
	go run $(MIGRATE_POINT)/migrate.go up

# Откат миграций
migrate-down:
	go run $(MIGRATE_POINT)/migrate.go down