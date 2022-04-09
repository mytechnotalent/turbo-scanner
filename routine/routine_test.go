// Terms Of Use
// ------------
// Do NOT use this on any computer you do not own or are not allowed to run this on.
// You may NEVER attempt to sell this, it is free and open source.
// The authors and publishers assume no responsibility.
// For educational purposes only.

package routine

import (
	"testing"
)

// Verify Routine scans TCP ports on the respective host.
func TestRoutineScansTCPPortsOnHost(t *testing.T) {
	// Params
	host := "localhost"
	ports := make(chan int, 1)
	results := make(chan int)
	var openports []int
	// Calls
	for i := 0; i < cap(ports); i++ {
		go Routine(&host, ports, results)
	}
	go func() {
		for i := 1; i <= 1; i++ {
			ports <- i
		}
	}()
	for i := 0; i < 1; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}
}
