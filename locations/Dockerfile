# Используем базовый образ
FROM golang:latest

# Устанавливаем рабочую директорию
WORKDIR /locations

# Копируем модульные файлы для зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код проекта в контейнер
COPY . .

# Собираем бинарный файл
RUN go build -o locations cmd/main.go

# Устанавливаем команду запуска контейнера
CMD ["./locations"]
