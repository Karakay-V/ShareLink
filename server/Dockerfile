# Базовий образ з Go
FROM golang:1.25-alpine

# Пакети для CGO
RUN apk add --no-cache build-base

# Робоча директорія всередині контейнера
WORKDIR /app

# Копіюємо модульні файли та завантажуємо залежності
COPY go.mod go.sum ./
RUN go mod download

# Копіюємо весь проект
COPY . .

# Збираємо бінарник
RUN go build -o server

# Створюємо папку для бази SQLite
RUN mkdir -p /app/data

# Відкриваємо порт
EXPOSE 8002

# Команда за замовчуванням
CMD ["./server"]
