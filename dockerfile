# Use the official Go image as the base image
FROM golang:1.16 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download and cache Go dependencies
RUN go mod download

# Copy the rest of the application code to the container
COPY . .

# Build the Go application
RUN go build -o app ./cmd/main.go

# Create a new stage for the final image
FROM golang:1.16

# Set the working directory inside the container
WORKDIR /app

# Copy the built executable from the previous stage
COPY --from=build /app/app .

# Expose the port your application is listening on
EXPOSE 8080

# Run the application
CMD ["./app"]