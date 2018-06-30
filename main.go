package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	port := ":8080"
	fmt.Println("server started on port", port)
	li, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}

		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()
	firstLine := true
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if firstLine {
			mux(conn, line)
			firstLine = false
		}
		if line == "" {
			// header is done
			break
		}
	}
}

func mux(conn net.Conn, reqHeader string) {
	reqMethod := strings.Fields(reqHeader)[0]
	reqURI := strings.Fields(reqHeader)[1]

	switch {
	case reqMethod == "GET" && reqURI == "/":
		index(conn)
	case reqMethod == "GET" && reqURI == "/pics/cow.jpg":
		cow(conn)
	case reqMethod == "GET" && reqURI == "apply":
		apply(conn)
	case reqMethod == "POST" && reqURI == "apply":
		applyPost(conn)
	default:
		notFound(conn)
	}
}

func writeResHeaders(conn net.Conn, contentType string) {
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprint(conn, "Content-Type:", contentType, "\r\n")
	io.WriteString(conn, "MyHeader: Loris\r\n")
	io.WriteString(conn, "\r\n")
}

func index(conn net.Conn) {
	defer conn.Close()

	body := `
		<!DOCTYPE html>
		<head>
			<meta charset="utf-8" />
			<title>Page Title</title>
		</head>
		<body>
			<h1>Holy cow this is low level!</h1>
			<img src="pics/cow.jpg">
		</body></html>`

	writeResHeaders(conn, "text/html")
	io.WriteString(conn, body)
}

func cow(conn net.Conn) {
	defer conn.Close()

	file, err := os.Open("cow.jpg")
	handleFileErr(err)
	defer file.Close()

	writeResHeaders(conn, "image/jpeg")
	io.Copy(conn, file)
}

func apply(conn net.Conn)     {}
func applyPost(conn net.Conn) {}

func notFound(conn net.Conn) {
	defer conn.Close()

	writeResHeaders(conn, "text/html")
	err := tpl.ExecuteTemplate(conn, "notFound.gohtml", nil)
	handleTemplateErr(err)
}

func handleTemplateErr(err error) {
	if err != nil {
		log.Fatalln("error executing template:", err)
	}
}

func handleFileErr(err error) {
	if err != nil {
		log.Fatalln("error opening file:", err)
	}
}
