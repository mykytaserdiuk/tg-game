# Этап сборки
FROM golang:1.22-alpine AS builder

# Установим рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download && go mod verify

# Копируем остальные исходные файлы
COPY . .

# Сборка проекта
RUN go build -v -o /app/main ./cmd/soap

# Этап запуска
FROM alpine:3.13

# Установим рабочую директорию
WORKDIR /app

# Копируем собранное приложение из предыдущего этапа
COPY --from=builder /app/main .

# Открываем порт 3000
EXPOSE 3000

# Указываем команду для запуска
CMD ["/app/main"]
