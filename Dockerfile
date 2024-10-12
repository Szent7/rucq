FROM golang:alpine AS builder
WORKDIR /build
ADD go.mod .
COPY . .
RUN go build -o main main.go 

FROM alpine
WORKDIR /build
COPY --from=builder /build/main /build/main
COPY --from=builder /build/images/ /build/images/
COPY --from=builder /build/templates/ /build/templates/

EXPOSE 10015
EXPOSE 10016
EXPOSE 10017

CMD ["./main"]