# Use the official Golang base image
FROM golang:latest

# Set the working directory in the container
WORKDIR /go/src/app

# Copy go.mod and go.sum to ensure dependencies are downloaded efficiently
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the current directory contents into the container
COPY . .

# Build the Go application
RUN go install -v ./...

# Set the entry point for the application
ENTRYPOINT ["app"]