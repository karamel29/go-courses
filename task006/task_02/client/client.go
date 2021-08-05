package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	var d net.Dialer
	fmt.Print("Enter \"exit\" to terminate the program!\n")

	for {
		fmt.Print("Enter your message: ")
		message, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		if message == "exit\n" {
			return
		}

		const port = 8081

		fmt.Printf("Sending message: %s; to port: %d\n", message[:len(message)-1], port)

		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		conn, err := d.DialContext(ctx, "tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		_, err = conn.Write([]byte(message))
		if err != nil {
			log.Fatal(err)
		}

		getMessage, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Message recieved: %s\n", getMessage[:len(getMessage)-1])
	}
}
