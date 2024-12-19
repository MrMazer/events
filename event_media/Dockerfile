# Используем базовый образ
FROM golang:latest

# Устанавливаем рабочую директорию
WORKDIR /event_media

# Копируем модульные файлы для зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код проекта в контейнер
COPY . .

# Собираем бинарный файл
RUN go build -o event_media cmd/main.go

# Устанавливаем команду запуска контейнера
CMD ["./event_media"]
