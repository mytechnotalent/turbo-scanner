package main

import (
	"fmt"
	"os"
	"runtime"
	"sort"

	"github.com/mytechnotalent/turbo-scanner/routine"
	"github.com/mytechnotalent/turbo-scanner/services"
	"github.com/mytechnotalent/turbo-scanner/setup"
)

func main() {
	if runtime.GOOS == "linux" {
		setup.LinuxSetup()
	}

	if runtime.GOOS == "windows" {
		if len(os.Args) != 2 {
			fmt.Println("usage: turbo-scanner_010.exe <host>")
			return
		}
	} else if runtime.GOOS == "darwin" {
		if len(os.Args) != 2 {
			fmt.Println("usage: ./turbo-scanner_010 <host>")
			return
		}
	} else if runtime.GOOS == "linux" {
		if len(os.Args) != 2 {
			fmt.Println("usage: sudo ./turbo-scanner_010 <host>")
			return
		}
	}

	host := os.Args[1]
	ports := make(chan int, 1000)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go routine.Routine(&host, ports, results)
	}

	go func() {
		for i := 1; i <= 65535; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 65535; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
		if runtime.GOOS == "windows" {
			fmt.Println(services.WinService(&port))
		} else if runtime.GOOS == "darwin" {
			fmt.Println(services.MacService(&port))
		} else if runtime.GOOS == "linux" {
			fmt.Println(services.LinuxService(&port))
		}
	}

}
