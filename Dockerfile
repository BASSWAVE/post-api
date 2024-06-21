# Stage 1: Build the Go app
FROM golang:1.22-alpine AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main ./cmd

# Stage 2: Create a small image to run the Go app
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the pre-built binary file from the previous stage
COPY --from=build /app/main .

# Command to run the executable
CMD ["./main"]
