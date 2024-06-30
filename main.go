package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/RoshanShrestha123/port-peek/port"
	"github.com/olekukonko/tablewriter"
)

func main() {

	startingPort := flag.Int("startPort", 1, "staring port address")
	ip := flag.String("ip", "127.0.0.1", "IP address to be scanned")
	endingPort := flag.Int("endPort", 5000, "Last port to scan")
	flag.Parse()

	var wg sync.WaitGroup
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"PORT", "NETWORK", "STATUS"})

	fmt.Println("---Searching for OPEN port---")
	for i := *startingPort; i < *endingPort; i++ {
		wg.Add(1)
		go port.ScanPort(*ip, i, &wg, table)
	}
	wg.Wait()
	table.Render()
	fmt.Println("Done")

}
