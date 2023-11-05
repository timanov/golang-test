# Используйте официальный образ Golang как базовый образ для сборки
FROM golang:1.17 AS build

# Установка рабочей директории
WORKDIR /app

# Копирование исходного кода в контейнер
COPY golang.go .

# Сборка бинарного файла myapp
RUN go build -o myapp golang.go

# Создание конечного образа
FROM debian:bullseye-slim

# Установка корневого сертификата
COPY ca-certificates.crt /etc/ssl/certs/

# Копирование бинарного файла из базового образа сборки
COPY --from=build /app/myapp /usr/local/bin/

# Запуск приложения при старте контейнера
CMD ["myapp"]
