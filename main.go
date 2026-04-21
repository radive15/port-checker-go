package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

var portsToCheck = []int{
	22,   // SSH
	80,   // HTTP
	443,  // HTTPS
	3306, // MySQL
	5432, // PostgreSQL
	6379, // Redis
}

func checkPort(host string, port int) bool {
	addr := net.JoinHostPort(host, fmt.Sprintf("%d", port))
	conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func main() {
	// ambil host dari argument pertama
	host := "localhost" // default jika tidak ada argument
	if len(os.Args) > 1 {
		host = os.Args[1]
	}

	fmt.Printf("Scanning %s...\n\n", host)

	openCount := 0
	for _, port := range portsToCheck {
		status := "CLOSED"
		if checkPort(host, port) {
			status = "OPEN  "
			openCount++
		}
		fmt.Printf("Port %-6d %s\n", port, status)
	}

	fmt.Printf("\n%d port terbuka dari %d yang dicek\n",
		openCount, len(portsToCheck))
}
