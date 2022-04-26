package main

import (
	"fmt"
	"runtime"
)

func main() {
	procs := runtime.GOMAXPROCS(0)
	cpunum := runtime.NumCPU()
	// out
	fmt.Println(procs)
	fmt.Println(cpunum)

}
