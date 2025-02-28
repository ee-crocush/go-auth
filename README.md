# Go Auth Service

**Go Auth** — это сервис аутентификации. 
Он предоставляет механизмы регистрации, входа, обновления токенов и выхода для пользователей. 
Это только пример, сервис не боевой

## Технологии

- **Go** — основной язык разработки.
- **gRPC** — механизм для создания RPC-сервисов.
- **PostgreSQL** — реляционная база данных для хранения данных.
- **JWT** — JSON Web Token для аутентификации пользователей.
- **Redis*** — для хранения сессий пользователей (реализация будет позже).

## Структура проекта

Проект реализует **чистую архитектуру**, разделяя код на несколько слоев, 
что улучшает тестируемость, поддерживаемость и расширяемость. 
Структура проекта следующая:

### Корень проекта
- **cmd/app** — точка входа. Содержит файл `main.go`, который вызывает 
- функцию `Run` из файла `internal/app/app.go`.

- **config** — конфигурации проекта, включая параметры для работы с базой данных, 
аутентификацией и другими сервисами.
- **internal** — реализация чистой архитектуры, включая следующие слои:
  - **app** — здесь реализуется запуск приложения.
  - **entity** — модели данных, представляющие бизнес-объекты, которые используются
  - **adapter** — взаимодействует с внешними системами и сервисами
  - **controller** — обрабатывает входящие запросы и передает их в слой `usecase` 
    для дальнейшей обработки. В этом проекте используется gRPC для обработки запросов.
  в других слоях, такие как пользователь.
  - **usecase** — бизнес-логика, которая инкапсулирует правила обработки данных и 
  выполнения операций, таких как аутентификация и управление сессиями.
  - **repo** — слой для работы с базой данных. В этом проекте используется PostgreSQL через 
  драйвер `lib/pq` для взаимодействия с хранилищем данных.
- **proto** - gprc контракты
- **gen** — сгенерированный код для gRPC.
- **pkg** — вспомогательные пакеты
- **migrations** — миграции базы данных.

**Для взаимодействия с redis описание в разработке*

## Установка

### 1. Клонирование репозитория

Для начала, клонируйте репозиторий:

```bash
git clone git@gitlab.dev-api.tech:cyberball/auth.git
```

### 2. Установка зависимостей

Перейдите в папку проекта и установите все зависимости:

```bash
go mod tidy
```

### 3. Настройка конфигурации

Создайте файл конфигурации: *config/config.yaml*. Пример:
```yaml
app:
  name: "app" # Имя приложения
  version: "0.0.1" # Версия приложения

grpc:
  port: 12345
  timeout: 30

logger:
  level: 'debug'

token:
  accessTTL: "30m"
  refreshTTL: "720h"

migrations:
  path: "path/to/migrations"
```
В корне проекта добавьте свой *.env* файл. Пример:
```text
MIGRATIONS_PATH=/custom/migrations

TOKEN_SECRET=SUPER_SECRET

# PostgreSQL тестовая конфигурация
PG_USER=admin
PG_PASSWORD=supersecurepassword
PG_HOST=db.example.com
PG_PORT=5432
PG_DBNAME=name
```

### 4. Запуск проекта

Перед запуском необходимо выполнить миграции базы данных:
```bash
go run cmd/migrate.go
```
или
```bash
make migrate-up
```
Затем запустите приложение:
```bash
go run cmd/main.go
```
или
```bash
make run
```

### 5. Запуск в Docker

Шаг 4 можно пропустить и запустить в Docker:
```bash
docker compose up --build
```

## Технология JWT (JSON Web Tokens)

Для аутентификации и авторизации в сервисе используется JWT. JWT позволяет безопасно передавать информацию между
клиентом и сервером в виде токенов.
- Access Token — короткоживущий токен, используемый для подтверждения аутентификации пользователя.
- Refresh Token — долгоживущий токен, используемый для получения нового access token, когда старый истекает.

### Пример работы с токенами
1. Пользователь регистрируется или входит в систему. 
2. Сервер генерирует два токена:
   - access_token (для доступа к защищённым маршрутам).
   - refresh_token (для обновления access token).
3. Токены передаются клиенту, который использует их для аутентификации при последующих запросах.

## Тестирование
Для запуска тестов используйте команду:

```bash
go test ./...
```
