# Строим GO приложуху
FROM golang:1.16-alpine as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY internal/Structures/go.mod .
COPY internal/XMLReader/go.mod .
COPY internal/Postgres/go.mod .

RUN go mod download

COPY . .

RUN go build -o main .

# Новая ступень сборки

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /app .

EXPOSE 8080
CMD ["./main"]