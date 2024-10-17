FROM golang:1.22.4-alpine

# Set the working directory
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o server .

# Add labels for metadata
LABEL Description="This is a Docker file"
LABEL Author="Youssef & Brahim"

# Expose the port the app runs on
EXPOSE 8080

RUN apk add --no-cache bash

# Command to run the application
CMD ["./server"]
