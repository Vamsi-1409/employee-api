FROM golang:1.26.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
