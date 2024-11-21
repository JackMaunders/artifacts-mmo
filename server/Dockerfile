FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY src ./src/

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/artifacts-mmo ./src/main.go

EXPOSE 8080

CMD ["./artifacts-mmo"]
