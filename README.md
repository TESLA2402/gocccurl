# gocccurl

gocccurl is a lightweight and versatile command-line tool inspired by curl, designed specifically for making HTTP requests tailored for RESTful APIs. With a focus on simplicity and ease of use, gocccurl supports common HTTP methods such as GET, POST, PUT, and DELETE, making it an essential companion for developers working with web services.

## Features

- **HTTP Methods:** Perform GET, POST, PUT, and DELETE requests effortlessly.
- **Custom Headers:** Include additional headers in your requests for enhanced flexibility.
- **Data Payload:** Easily send data in the request body, supporting various content types.
- **Verbose Mode:** Enable verbose mode to see detailed information about the request and response.

## Usage

Get started with gocccurl by providing a URL and specifying the desired HTTP method. Additional options are available to customize headers and include data in the request body.

The following options are supported:

- `-v`: verbose mode
- `-X <method>`: allows us to specify the HTTP method
- `-d`: allows us to pass data to the server
- `-H`: allows us to include the headers

## Installation

Clone the repository and use the provided `build.sh` script to build mycurl. Ensure that you have Go installed on your machine.

```bash
git clone https://github.com/TESLA2402/gocccurl.git
cd curl
chmod +x build.sh
./build.sh

```

### Examples:

```bash
# If build script implemented

# Simple GET request
mycurl http://eu.httpbin.org/get

# DELETE request
mycurl -X DELETE http://eu.httpbin.org/delete

# POST request with JSON payload
mycurl -X POST http://eu.httpbin.org/post -d '{"key": "value"}' -H "Content-Type: application/json"

# PUT request with JSON payload
mycurl -X PUT http://eu.httpbin.org/put -d '{"key2": "value2"}' -H "Content-Type: application/json"

 OR

# Simple GET request
go run main.go http://eu.httpbin.org/get

# DELETE request
go run main.go -X DELETE http://eu.httpbin.org/delete

# POST request with JSON payload
go run main.go -X POST http://eu.httpbin.org/post -d '{"key": "value"}' -H "Content-Type: application/json"

# PUT request with JSON payload
go run main.go -X PUT http://eu.httpbin.org/put -d '{"key2": "value2"}' -H "Content-Type: application/json"
```
