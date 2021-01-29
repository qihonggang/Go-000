package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:10000")
	defer conn.Close()
	if err != nil {
		fmt.Println("Error dialing", err)
	}
	for {
		var content string
		fmt.Print("input message: ")
		_, _ = fmt.Scan(&content)
		contentList := []byte(content)
		_, err = conn.Write(contentList)
		if err != nil {
			fmt.Println(err)
		}
	}
}
