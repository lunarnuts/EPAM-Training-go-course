package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	port := 8081
	fmt.Printf("Connecting to port: %d\n", port)
	conn, err := net.Dial("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Message to send: ")
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Print(err)
			continue
		}
		_, err = conn.Write([]byte(message))
		if err != nil {
			log.Print(err)
			return
		}
		if strings.Trim(message, "\n") == "exit" {
			fmt.Println("exiting...")
			break
		}
		message, err = bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("Message from server: %s", message)
	}
}
