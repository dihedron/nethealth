package start

import (
	"time"

	"github.com/dihedron/nethealth/cmd/nethealth/command/base"
)

type Duration time.Duration

type Topology struct {
	Frequency Duration
	Nodes     []string
}

type Start struct {
	base.Command

	Topology Topology `short:"t" long:"topology" description:"The network endpoints to ping." required:"yes"`
}

func (cmd *Start) Execute(args []string) (err error) {
	logger := cmd.InitLogger(true)

	logger.Debug("starting...")
	return nil
}
