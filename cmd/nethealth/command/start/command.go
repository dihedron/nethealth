package start

import (
	"fmt"
	"sync"
	"time"

	"github.com/dihedron/nethealth/cmd/nethealth/command/base"
	probing "github.com/prometheus-community/pro-bing"
)

type Duration time.Duration

type Topology struct {
	Frequency Duration
	Nodes     []string
}

type Start struct {
	base.Command

	Topology Topology `short:"t" long:"topology" description:"The network endpoints to ping." optional:"yes"`
}

func (cmd *Start) Execute(args []string) (err error) {
	logger := cmd.InitLogger(true)

	logger.Debug("starting...")

	var wg sync.WaitGroup

	for _, address := range []string{
		"www.google.com",
		// "www.repubblica.it",
		"www.bbc.co.uk",
	} {

		wg.Add(1)

		go func(address string) {
			pinger, err := probing.NewPinger(address)
			if err != nil {
				panic(err)
			}
			pinger.Count = 10
			pinger.SetPrivileged(true)
			pinger.Interval = 1 * time.Millisecond
			err = pinger.Run() // Blocks until finished.
			if err != nil {
				panic(err)
			}
			stats := pinger.Statistics() // get send/receive/duplicate/rtt stats
			fmt.Printf("stats to %s: %v\n", address, stats)
			wg.Done()
		}(address)
	}

	wg.Wait()
	fmt.Printf("done!\n")
	return nil
}
