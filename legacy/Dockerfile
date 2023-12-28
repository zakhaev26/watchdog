# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's working directory
COPY . .

# Download and install any dependencies
RUN go get -u github.com/gorilla/mux

# Build the Go application
RUN go build -o main .

# Expose the port on which the application will run
EXPOSE 8080

# Command to run the application
CMD ["./main"]

