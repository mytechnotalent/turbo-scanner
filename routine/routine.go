// Terms Of Use
// ------------
// Do NOT use this on any computer you do not own or are not allowed to run this on.
// You may NEVER attempt to sell this, it is free and open source.
// The authors and publishers assume no responsibility.
// For educational purposes only.

// Package routine provides port scanning functionality.
package routine

import (
	"fmt"
	"net"
)

// Routine takes a host and scans 65,535 TCP ports on the respective host.
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
