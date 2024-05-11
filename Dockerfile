FROM golang:1.22.3-alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server

EXPOSE 8080
CMD ["/app/main"]