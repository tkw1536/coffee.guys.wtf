FROM golang as builder

WORKDIR /app/
ADD *.go /app/
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main /app/main.go

FROM alpine
EXPOSE 8080

WORKDIR /app/
COPY --from=builder /app/main /app/main
ADD static /app/static
ENTRYPOINT ["/app/main"]