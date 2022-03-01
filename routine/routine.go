// Package routine provides port scanning functionality.
package routine

import (
	"fmt"
	"net"
)

// Routine takes a host and scans 65,535 tcp ports on the respective host.
func Routine(host *string, ports, results chan int) {
	for port := range ports {
		address := fmt.Sprintf("%s:%d", *host, port)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- port
	}
}
