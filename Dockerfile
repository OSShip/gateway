FROM golang:1.26-alpine AS builder
WORKDIR /app
RUN go install github.com/swaggo/swag/cmd/swag@v1.16.6
COPY utils/ /app/utils/
COPY services/gateway/ /app/services/gateway/
WORKDIR /app/services/gateway
RUN swag init -g main.go --parseInternal --output docs
RUN go mod download && CGO_ENABLED=0 go build -o /gateway .

FROM alpine:3.20
RUN apk add --no-cache ca-certificates wget
COPY --from=builder /gateway /gateway
EXPOSE 8080
CMD ["/gateway"]
