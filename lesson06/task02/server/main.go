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
	port := 8081
	fmt.Printf("Listening on port: %d\n", port)
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print(err)
			break
		}
		go func(c net.Conn) {
			defer c.Close()
			for {
				message, err := bufio.NewReader(c).ReadString('\n')
				if err != nil {
					log.Print(err)
					return
				}
				fmt.Printf("Message received: %s", message)
				if strings.Trim(message, "\n") == "exit" {
					fmt.Println("Client closed")
					break
				}
				num, err := strconv.Atoi(strings.TrimSpace(message))
				if err != nil {
					_, er := c.Write([]byte(strings.ToUpper(message) + "\n"))
					if er != nil {
						log.Print(err)
						return
					}
					continue
				}
				num *= 2
				_, err = c.Write([]byte(strconv.Itoa(num) + "\n"))
				if err != nil {
					log.Print(err)
					return
				}
			}
		}(conn)
	}

}
