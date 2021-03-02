# Build the app
FROM golang as builder

# build the app
WORKDIR /app/
ADD . /app/
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main /app/main.go

# Create a new user 'www-data'
FROM alpine as permission

ENV USER=www-data
ENV UID=82
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"


# add it into a scratch image
FROM scratch
WORKDIR /app/

# add the user
COPY --from=permission /etc/passwd /etc/passwd
COPY --from=permission /etc/group /etc/group

# add the app and static files
COPY --from=builder /app/main /app/main

# and set the entry command
EXPOSE 8080
USER www-data:www-data
ENTRYPOINT ["/app/main"]