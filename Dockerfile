# Этап сборки приложения
FROM golang:1.20-alpine AS builder

WORKDIR /app

# Копируем файлы go.mod и go.sum для скачивания зависимостей (если есть)
COPY go.mod go.sum ./

# Загружаем зависимости (кэширование)
RUN go mod download

# Копируем весь исходный код проекта в контейнер
COPY . .

# Собираем приложение в бинарник 'app'
RUN go build -o app ./cmd/url-shortener

# Финальный образ — минимальный Alpine Linux с нашим приложением
FROM alpine:latest

WORKDIR /app

# Копируем собранный бинарник из предыдущего этапа
COPY --from=builder /app/app .

# Указываем команду запуска при старте контейнера
CMD ["./app"]
