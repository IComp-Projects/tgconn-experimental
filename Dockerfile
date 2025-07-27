FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o tgconn cmd/*.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/tgconn .

ENTRYPOINT ["/app/tgconn"]
