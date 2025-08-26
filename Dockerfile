FROM golang as builder

WORKDIR /app/

COPY . .

RUN CGO_ENABLED=0 go build -o inventory-service /app/cmd/grpc/main.go
RUN CGO_ENABLED=0 go build -o inventory-service-worker /app/cmd/worker/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/ /app/
COPY --from=builder /app/inventory-service-worker /app/

#EXPOSE 5002
