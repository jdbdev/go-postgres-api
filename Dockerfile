# Golang image:
FROM golang:1.24-alpine3.21

# Working directory:
WORKDIR /app

# Copy source code to /app:
COPY . .

# Dowload & install dependencies:
RUN go get -d -v ./...

# Build the Go application (name is in go.mod:
RUN go build -o api .

# Expose port in Docker:
EXPOSE 8000

# Run the executable:
CMD ["./api"]