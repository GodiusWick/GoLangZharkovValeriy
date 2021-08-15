# Строим GO приложуху
FROM golang:latest as builder

WORKDIR /app

COPY . .

RUN go build -o main .

# Новая ступень сборки

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]