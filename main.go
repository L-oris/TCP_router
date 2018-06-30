package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/L-oris/tcpMux/routes"
)

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
		routes.Index(conn)
	case reqMethod == "GET" && reqURI == "/pics/cow.jpg":
		routes.Cow(conn)
	case reqMethod == "GET" && reqURI == "/about":
		routes.About(conn)
	default:
		routes.NotFound(conn)
	}
}
