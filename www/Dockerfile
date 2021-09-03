# Строим GO приложуху
FROM golang:1.16-alpine as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

COPY postgres/go.mod .
COPY structures/go.mod .
COPY xmlreader/go.mod .

RUN go mod download

COPY . .

RUN go build -o main .

# Новая ступень сборки

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /app .

EXPOSE 8080
CMD ["./main"]