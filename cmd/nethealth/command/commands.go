package command

import (
	"github.com/dihedron/nethealth/cmd/nethealth/command/start"
	"github.com/dihedron/nethealth/cmd/nethealth/command/version"
)

// Commands is the set of root command groups.
type Commands struct {
	// Version prints the application version and exits.
	Version version.Version `command:"version" alias:"ver" alias:"v" description:"Print the command version and exit."`
	// Start starts the network health monitor.
	Start start.Start `command:"start" alias:"s" description:"Start the network health monitor."`
}
