package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	
	"github.com/akamensky/argparse"
)

func main() {
    parser := argparse.NewParser("tcpget", "receive data from tcp sockets")
    addr := parser.String("a", "address", &argparse.Options{Required: true, Help: "ip address of the server"})
    port := parser.String("p", "port", &argparse.Options{Required: true, Help: "the tcp port to connect to"})
    if err := parser.Parse(os.Args); err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}
	handle_connections(*addr, *port)
}

func handle_connections(ip, port string) {
    conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", ip, port))
    if err != nil {
        fmt.Printf("ERROR: %s\n", err)
        os.Exit(1)
    }
    buf := bufio.NewReader(conn)
    for {
        data, err := buf.ReadString('\n')
        if err != nil {
            fmt.Printf("ERROR: %s", err)
            os.Exit(1)
        }
        fmt.Println(data)
    }
}