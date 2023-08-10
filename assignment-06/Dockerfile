FROM golang:1.20-alpine AS builder
WORKDIR /build
COPY . ./
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o app .

FROM alpine:latest
COPY --from=builder /build/app ./
ENTRYPOINT ["/app"]