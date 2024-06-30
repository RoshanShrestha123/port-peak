package port

import (
	"net"
	"strconv"
	"sync"

	"github.com/olekukonko/tablewriter"
)

func ScanPort(address string, port int, wg *sync.WaitGroup, table *tablewriter.Table) {
	_, err := net.Dial("tcp", address+":"+strconv.Itoa(port))

	defer wg.Done()

	if err != nil {
		return
	}

	table.Append([]string{strconv.Itoa(port), "TCP", "OPEN"})

}
