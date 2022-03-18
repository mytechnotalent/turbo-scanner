// Terms Of Use
// ------------
// Do NOT use this on any computer you do not own or are not allowed to run this on.
// You may NEVER attempt to sell this, it is free and open source.
// The authors and publishers assume no responsibility.
// For educational purposes only.

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
