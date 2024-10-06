FROM golang:alpine AS builder
WORKDIR /build
ADD go.mod .
COPY . .
RUN go build -o main main.go client.go hub.go

FROM alpine
WORKDIR /build
COPY --from=builder /build/main /build/main
COPY --from=builder /build/home.html /build/home.html
EXPOSE 10016
CMD ["./main"]