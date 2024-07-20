# Use the official Golang image as the base image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o learn1 .

# Use a minimal image for the final build


# Expose the application port (optional)
EXPOSE 8080

# Command to run the executable
CMD ["./learn1"]
