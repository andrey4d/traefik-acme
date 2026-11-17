FROM golang:1.26-alpine as builder

ARG GOOS=linux
ARG GOARCH=amd64

WORKDIR /app
COPY . /app
RUN go mod download
RUN go build -o build/traefik-acme .

# Use a minimal image for the final container
FROM scratch
COPY --from=builder /app/build/traefik-acme /traefik-acme
CMD ["/traefik-acme", "get"]