// Package services provides service specific OS functionality.
package services

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// WinService takes a port to obtain the running service on it.
// It returns the Windows service name and respective details.
func WinService(port *int) string {
	arg0 := "Get-Process -Id (Get-NetTCPConnection -LocalPort "
	arg1 := strconv.Itoa(*port)
	arg2 := ").OwningProcess"
	args := arg0 + arg1 + arg2
	out, _ := exec.Command("powershell", "/C", args).Output()
	trimmedOut := strings.TrimSpace(string(out))
	trimmedOut += "\n"
	return string(trimmedOut)
}

// LinuxService takes a port to obtain the running service on it.
// It returns the Debian-based Linux service name and respective details.
func LinuxService(port *int) string {
	arg0 := "netstat -tlnp | grep "
	arg1 := strconv.Itoa(*port)
	args := arg0 + arg1
	out, err := exec.Command("sh", "-c", args).Output()
	if err != nil {
		log.Fatal("sudo apt install net-tools")
	}
	return string(out)
}

// MacService takes a port to obtain the running service on it.
// It returns the Mac service name and respective details.
func MacService(port *int) string {
	arg0 := "lsof -i :"
	arg1 := strconv.Itoa(*port)
	args := arg0 + arg1
	out, _ := exec.Command("sh", "-c", args).Output()
	return string(out)
}
