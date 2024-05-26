FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /app/main .

FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 1232
CMD ["app/main"]