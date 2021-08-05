package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func main() {

	const port = 8081

	fmt.Printf("Launching server on port: %d \n\n", port)

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go func(c net.Conn) {
			defer c.Close()

			message, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
				log.Print(err)
				return
			}

			fmt.Printf("Message received: %s \n", message)

			message = message[:len(message)-1]
			val, err := strconv.Atoi(message)
			var messageBack string
			if err != nil {
				messageBack = strings.ToUpper(message)
			} else {
				messageBack = strconv.Itoa(val * 2)
			}

			_, err = c.Write([]byte(messageBack + "\n"))
			if err != nil {
				log.Print(err)
				return
			}
		}(conn)
	}
}
