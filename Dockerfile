# Use the official Golang image to create a build artifact.
FROM golang:1.23 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files for dependency management
COPY go.mod go.sum ./

# Copy go.mod and go.sum files for dependency management
RUN go mod download

# Copy the source code into the container
COPY . ./

# Build the Go app
RUN go build -o ../ ./cmd

