package main // import "github.com/secretbox-io/secretbox"
import (
	"github.com/secretbox-io/secretbox/cmd"
	log "github.com/sirupsen/logrus"
)

// Version is the tagged build number
var Version string

// CommitHash is the current commit for the build
var CommitHash string

// BuildTime is the build time
var BuildTime string

func main() {
	cmd.Version = Version
	cmd.CommitHash = CommitHash
	cmd.BuildTime = BuildTime

	log.SetFormatter(&log.TextFormatter{
		DisableTimestamp:       true,
		DisableLevelTruncation: true,
	})

	cmd.Execute()
}
