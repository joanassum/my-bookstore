# bookstore-books

Microservice for storing and retrieving information of books.

## Installation

1. Install Go and go-swagger.

2. Generate code for server by running `swagger generate server -A bookstore-books -f ./swagger.yml`

3. Build by running `go build -o bookstore-books-server ./cmd/bookstore-books-server`, or for docker `docker build -t bookstore-books:dev .`

4. Run `./cmd/bookstore-books-server --port 8080`, or for docker `docker run -it --rm  -p 8080:8080  bookstore-books:dev`. 

## Architecture

Use clean architecture with reference to [](https://github.com/bxcodec/go-clean-arch)