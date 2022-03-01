// Package setup provides OS specific setup functionality.
package setup

import "os/exec"

// LinuxSetup installs the net-tools library to ensure netstat has the ability to run.
// It returns any necessary info to stdin from the result of the OS install.
func LinuxSetup() string {
	arg0 := "sudo apt install net-tools"
	out, _ := exec.Command("sh", "-c", arg0).Output()
	return string(out)
}
