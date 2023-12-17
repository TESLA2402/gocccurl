package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var verbose bool
var method string
var data string
var contentType string

func init() {
	flag.Usage = func() {
		fmt.Println("Usage: go run main.go -X <method> [options] <URL> -d <data> -H <contentType>")
		flag.PrintDefaults()
	}

	flag.BoolVar(&verbose, "v", false, "Enable verbose mode")
	flag.StringVar(&method, "X", "GET", "Specify the HTTP method (GET by default)")
	flag.StringVar(&data, "d", "", "Specify the data to be sent in the request body")
	flag.StringVar(&contentType, "H", "", "Specify additional headers")
	flag.Parse()
}

func main() {
	if len(flag.Args()) < 1 {
		flag.Usage()
		return
	}

	url := flag.Arg(0)
	var req *http.Request
	var err error

	if method == "POST" || method == "PUT" {
		data := flag.Arg(2)
		fmt.Println(data)
		req, err = http.NewRequest(method, url, strings.NewReader(data))
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		contentType := strings.Split(flag.Arg((4)), "Content-Type: ")
		if contentType != nil {
			req.Header.Set("Content-Type", contentType[len(contentType)-1])
		} else {
			req.Header.Set("Content-Type", "application/json")
		}
	} else {
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}
	}

	// Add additional headers if specified
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	// Print the request data if in verbose mode
	if verbose {
		fmt.Printf("> %s %s HTTP/1.1\n", method, url)
		req.Header.Write(os.Stdout)
		if method == "POST" || method == "PUT" {
			fmt.Println("\r\n" + data)
		} else {
			fmt.Println()
		}
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	if verbose {
		resp.Header.Write(os.Stdout)
		fmt.Println()
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Print the response data
	fmt.Println(string(body))
}
