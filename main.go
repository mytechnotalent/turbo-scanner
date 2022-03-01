package main

import (
	"fmt"
	"runtime"
	"sort"

	"github/mytechnotalent/turbo-scanner/routine"
)

func main() {
	if runtime.GOOS == "linux" {
		linuxSetup()
	}

	ports := make(chan int, 1000)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go routine.Routine(ports, results)
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
			fmt.Println(winService(&port))
		} else if runtime.GOOS == "darwin" {
			fmt.Println(macService(&port))
		} else if runtime.GOOS == "linux" {
			fmt.Println(linuxService(&port))
		}
	}

}
