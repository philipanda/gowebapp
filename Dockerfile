FROM golang:alpine AS builder
WORKDIR /go/src/GoWebApp
COPY . .
RUN go mod download
RUN go build

FROM alpine:latest
WORKDIR /root/
COPY . .
COPY --from=builder /go/src/GoWebApp/GoWebApp .
CMD ["./GoWebApp"]