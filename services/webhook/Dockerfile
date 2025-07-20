FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /webhook

FROM scratch

WORKDIR /app

COPY --from=builder /webhook .

EXPOSE 5555

CMD [ "/app/webhook" ]
