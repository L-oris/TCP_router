package routes

import (
	"fmt"
	"html/template"
	"io"
	"net"
	"os"

	"github.com/L-oris/tcpMux/people"
	"github.com/L-oris/tcpMux/utils"
)

func writeResHeaders(conn net.Conn, contentType string) {
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprint(conn, "Content-Type:", contentType, "\r\n")
	io.WriteString(conn, "MyHeader: Loris\r\n")
	io.WriteString(conn, "\r\n")
}

// Index page
func Index(conn net.Conn) {
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

// Cow image
func Cow(conn net.Conn) {
	defer conn.Close()

	file, err := os.Open("assets/cow.jpg")
	utils.HandleFileErr(err)
	defer file.Close()

	writeResHeaders(conn, "image/jpeg")
	io.Copy(conn, file)
}

// About page
func About(conn net.Conn) {
	defer conn.Close()

	writeResHeaders(conn, "text/html")
	tpl := template.Must(template.ParseFiles("templates/about.gohtml"))
	err := tpl.Execute(conn, people.GeneratePeople())
	utils.HandleTemplateErr(err)
}

// NotFound page
func NotFound(conn net.Conn) {
	defer conn.Close()

	writeResHeaders(conn, "text/html")
	tpl := template.Must(template.ParseFiles("templates/notFound.gohtml"))
	err := tpl.Execute(conn, nil)
	utils.HandleTemplateErr(err)
}
