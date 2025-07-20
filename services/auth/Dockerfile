FROM golang:1.24-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /auth

FROM scratch

WORKDIR /app

COPY --from=builder /auth .

EXPOSE 4444

CMD [ /app/auth ]

