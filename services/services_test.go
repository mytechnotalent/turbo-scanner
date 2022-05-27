// Terms Of Use
// ------------
// Do NOT use this on any computer you do not own or are not allowed to run this on.
// You may NEVER attempt to sell this, it is free and open source.
// The authors and publishers assume no responsibility.
// For educational purposes only.

package services

import (
	"testing"
)

// Verify WinService obtains running service.
func TestWinServiceObtainsRunningService(t *testing.T) {
	// Params
	port := 1
	_, err := WinService(&port)
	// Asserts
	if err == nil {
		t.Error(err)
	}
}

// Verify LinuxService obtains running service.
func TestLinuxServiceObtainsRunningService(t *testing.T) {
	// Params
	port := 1
	_, err := LinuxService(&port)
	// Asserts
	if err == nil {
		t.Error(err)
	}
}

// Verify MACService obtains running service.
func TestMACServiceObtainsRunningService(t *testing.T) {
	// Params
	port := 1
	_, err := MACService(&port)
	// Asserts
	if err == nil {
		t.Error(err)
	}
}
