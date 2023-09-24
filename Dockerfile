# Use the official Golang base image
FROM golang:latest

# Set the working directory in the container
WORKDIR /go/src/app

# Copy the current directory contents into the container at /go/src/app
COPY . /go/src/app

# Build the Go application
RUN go get -d -v ./...
RUN go install -v ./...

# Set the entry point for the application
ENTRYPOINT ["app"]
