package main

import (
	"bufio"
	"log"
	"net"
)

func handleReadConn(conn net.Conn, ch chan <- string) {
	defer conn.Close()
	defer close(ch)
	rd := bufio.NewReader(conn)
	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			log.Printf("read error: %v\n", err)
			return
		}
		ch <- string(line)
	}
}

func handleWriteConn(conn net.Conn, ch <- chan string) {
	wr := bufio.NewWriter(conn)
	for {
		msg, ok := <- ch
		if !ok {
			break
		}
		wr.WriteString("msg: ")
		wr.WriteString(msg)
		wr.Flush()
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:10000")
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept error: %v\n", err)
			continue
		}
		ch := make(chan string, 10)
		// 读消息
		go handleReadConn(conn, ch)
		// 写消息
		go handleWriteConn(conn, ch)
	}
}