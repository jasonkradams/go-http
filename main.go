package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

var (
	fileName   string
	headers    headerList
	statusCode int
)

type headerList []string

func (h *headerList) String() string {
	return strings.Join(*h, ", ")
}

func (h *headerList) Set(value string) error {
	*h = append(*h, value)
	return nil
}

func main() {
	flag.StringVar(&fileName, "file", "", "File to serve as /index.html")
	flag.Var(&headers, "header", "Custom headers to add to the response (repeatable)")
	flag.IntVar(&statusCode, "status", 200, "Status response code (200, 404, 500)")
	flag.Parse()

	if fileName == "" {
		fmt.Println("Error: -file argument is required")
		return
	}

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fmt.Printf("Error: file %s does not exist\n", fileName)
		return
	}

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Starting server at port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading request:", err)
		return
	}

	requestLine = strings.TrimSpace(requestLine)
	method, path, version, err := parseRequestLine(requestLine)
	if err != nil || method != "GET" || !strings.HasPrefix(version, "HTTP/") {
		sendResponse(conn, "HTTP/1.1 400 Bad Request", "", "")
		return
	}

	if path == "/" {
		path = "/index.html"
	}

	if path != "/index.html" {
		sendResponse(conn, "HTTP/1.1 404 Not Found", "", "")
		return
	}

	serveFile(conn, fileName)
}

func parseRequestLine(requestLine string) (method, path, version string, err error) {
	parts := strings.Split(requestLine, " ")
	if len(parts) != 3 {
		err = fmt.Errorf("invalid request line: %s", requestLine)
		return
	}
	method, path, version = parts[0], parts[1], parts[2]
	return
}

func sendResponse(conn net.Conn, statusLine, headers, body string) {
	response := statusLine + "\r\n" + headers + "\r\n" + body
	conn.Write([]byte(response))
}

func serveFile(conn net.Conn, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		sendResponse(conn, "HTTP/1.1 404 Not Found", "", "")
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		sendResponse(conn, "HTTP/1.1 500 Internal Server Error", "", "")
		return
	}
	fileSize := fileInfo.Size()

	headers := buildHeaders(fileSize)

	statusLine := getStatusLine()
	sendResponse(conn, statusLine, headers+"\r\n", "")

	io.Copy(conn, file)
}

func buildHeaders(fileSize int64) string {
	var headerStrings []string
	for _, header := range headers {
		headerStrings = append(headerStrings, header)
	}
	headerStrings = append(headerStrings, fmt.Sprintf("Content-Length: %d", fileSize))
	return strings.Join(headerStrings, "\r\n")
}

func getStatusLine() string {
	switch statusCode {
	case 200:
		return "HTTP/1.0 200 OK"
	case 404:
		return "HTTP/1.0 404 Not Found"
	case 500:
		return "HTTP/1.0 500 Internal Server Error"
	default:
		return "HTTP/1.0 500 Internal Server Error"
	}
}
