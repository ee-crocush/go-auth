# Этап сборки
FROM golang:1.23 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы для управления зависимостями
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download

# Копируем всё остальное
COPY . .

# Сборка бинарного файла приложения
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-auth ./cmd/app/main.go

# Сборка миграционного инструмента
RUN CGO_ENABLED=0 GOOS=linux go build -o /migrate ./cmd/migrate/migrate.go

# Финальный минимальный образ
FROM debian:bullseye-slim

# Устанавливаем минимальные зависимости для работы приложения
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем из builder-этапа бинарный файл приложения и миграционный инструмент
COPY --from=builder /go-auth /app/go-auth
COPY --from=builder /migrate /app/migrate

# Копируем конфигурационный файл и директорию с миграциями
COPY config/config.yaml /app/config/config.yaml
COPY migrations /app/migrations

# Указываем порт, на котором работает приложение
EXPOSE 60051

# Запускаем выполнение миграций перед запуском приложения
CMD ["/bin/sh", "-c", "/app/migrate up && /app/go-auth"]