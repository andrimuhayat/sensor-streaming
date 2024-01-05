FROM golang:latest as builder
COPY . /app/go

WORKDIR /app/go
ENV GO111MODULE=on
RUN go mod tidy
#RUN swag init
RUN CGO_ENABLED=0 GOOS=linux go build -o go
#second stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/go .
EXPOSE 8081

CMD ["./go"]