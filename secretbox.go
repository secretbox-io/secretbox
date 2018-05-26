package main // import "github.com/secretbox-io/secretbox"
import (
	"github.com/secretbox-io/secretbox/cmd"
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

	cmd.Execute()
}
