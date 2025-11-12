# Configuration Guide

Приложение поддерживает загрузку конфигурации из файла и переменных окружения. **Переменные окружения имеют приоритет над значениями из файла конфигурации.**

## Порядок загрузки конфигурации

1. **Файл конфигурации** (`config/default.yml`) - загружается первым как значения по умолчанию
2. **Переменные окружения** - перезаписывают значения из файла

## Поддерживаемые переменные окружения

### Общие настройки

| Переменная | Описание | Пример |
|------------|----------|--------|
| `PORT` | Порт для HTTP сервера | `8081` |

### Настройки PostgreSQL

| Переменная | Описание | Пример |
|------------|----------|--------|
| `POSTGRES_HOST` | Хост базы данных | `localhost` |
| `POSTGRES_PORT` | Порт базы данных | `5432` |
| `POSTGRES_USER` | Имя пользователя | `aleksandr` |
| `POSTGRES_PASSWORD` | Пароль | `mypassword` |
| `POSTGRES_DB` | Имя базы данных | `NIR` |
| `DB_CONNECT_RETRY` | Количество попыток подключения | `3` |
| `DB_POOL_SIZE` | Размер пула соединений | `10` |
| `SSL_MODE` | Режим SSL | `disable`, `require`, `verify-full` |

## Примеры использования

### Локальная разработка

Используйте файл `config/default.yml`:

```yaml
postgres:
  database: NIR
  user: aleksandr
  password: ""
  host: localhost
  port: "5432"
  retries: 3
  pool_size: 10
  ssl_mode: disable

port: "8081"
```

### Docker / Production

Используйте переменные окружения:

```bash
export POSTGRES_HOST=postgres.example.com
export POSTGRES_PORT=5432
export POSTGRES_USER=prod_user
export POSTGRES_PASSWORD=secure_password
export POSTGRES_DB=vkr_production
export DB_POOL_SIZE=20
export SSL_MODE=require
export PORT=8080
```

### Docker Compose

```yaml
services:
  vkr-backend:
    environment:
      - POSTGRES_HOST=postgres
      - POSTGRES_PORT=5432
      - POSTGRES_USER=vkr_user
      - POSTGRES_PASSWORD=${POSTGRES_PASSWORD}
      - POSTGRES_DB=vkr_db
      - PORT=8081
```

### Kubernetes

```yaml
env:
  - name: POSTGRES_HOST
    value: "postgres-service"
  - name: POSTGRES_PORT
    value: "5432"
  - name: POSTGRES_USER
    valueFrom:
      secretKeyRef:
        name: postgres-credentials
        key: username
  - name: POSTGRES_PASSWORD
    valueFrom:
      secretKeyRef:
        name: postgres-credentials
        key: password
  - name: POSTGRES_DB
    value: "vkr_production"
  - name: PORT
    value: "8081"
```

## Приоритет значений

Если установлена переменная окружения, она всегда перезапишет значение из файла конфигурации:

```bash
# В config/default.yml: port: "8081"
# В переменных окружения: PORT=9090
# Результат: приложение запустится на порту 9090
```

## Проверка конфигурации

Для проверки загруженной конфигурации можно использовать логирование или добавить endpoint для отображения конфигурации (без чувствительных данных).

