// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 222.

// Clock is a TCP server that periodically writes the time.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	area := Beijing
	addr := c.LocalAddr().String()
	port := addr[strings.LastIndex(addr, ":")+1:]
	fmt.Printf("Conn is port: %s\n", port)
	area = portToArea[port]
	for {
		location, err := time.LoadLocation(areaToTimeZoo[area])
		if err != nil {
			log.Fatal("set time zone failed: " + err.Error())
		}
		_, err = io.WriteString(c, area+" time is: "+time.Now().In(location).Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var inputPort string
	flag.StringVar(&inputPort, "port", "8010", "clock port")
	flag.Parse()
	fmt.Printf("New clock in port: %s\n", inputPort)
	listener, err := net.Listen("tcp", "localhost:"+inputPort)
	if err != nil {
		log.Fatal(err)
	}
	//!+
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
	//!-
}
