FROM golang:alpine AS Builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Install basic packages
RUN apk add \
    gcc \
    g++

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go mod download

# Build image
RUN go build .

FROM alpine:latest AS Runner

WORKDIR /app

COPY templates ./templates
COPY --from=Builder /app/misso /app/app

# This container exposes port 8080 to the outside world
EXPOSE 8080/tcp

ENV MODE=prod

# Run the executable
CMD ["/app/app"]
