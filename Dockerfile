# Choose whatever you want, version >= 1.16
FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

RUN go install github.com/air-verse/air@latest

COPY . .
COPY ./.env /

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]