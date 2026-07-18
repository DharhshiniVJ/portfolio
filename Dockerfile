FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy the go.mod file
COPY go.mod ./

# Download dependencies (though we only have standard library right now)
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go application
RUN go build -o portfolio main.go

# Use a lightweight alpine image for the final stage
FROM alpine:latest

WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /app/portfolio .

# Copy the templates and static folders
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static

# Set the default port
ENV PORT=8081

# Expose the port
EXPOSE 8081

# Run the executable
CMD ["./portfolio"]
