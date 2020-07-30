FROM golang:1.14

# Set the Current Working Directory inside the container
WORKDIR /app/bookstore-books

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN CGO_ENABLED=0 go build -o ./out/bookstore-books ./cmd/bookstore-books-server/main.go

RUN chmod +x ./out/bookstore-books

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["./out/bookstore-books", "--port", "8080", "--host", "0.0.0.0"]