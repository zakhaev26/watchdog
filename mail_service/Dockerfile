# Use the official Go image as a base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container
COPY . /app

# Smtp Env Vars

ENV EMAIL_SENDER_NAME=${EMAIL_SENDER_NAME}
ENV EMAIL_SENDER_ADDRESS=${EMAIL_SENDER_ADDRESS}
ENV EMAIL_SENDER_PASSWORD=${EMAIL_SENDER_PASSWORD}

# Build the Go application
RUN go build main.go

# Expose the port the app runs on
EXPOSE 6969

# Command to run the executable
CMD ["./main"]