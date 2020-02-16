package version

import (
	"fmt"

	"github.com/jetstack/preflight/pkg/version"
)

// Injected at build time.

// Commit is the commit hash of the build.
var Commit string

// BuildDate is the date it was built.
var BuildDate string

// GoVersion is the go version that was used to compile this.
var GoVersion string

// Platform is the target platform this was compiled for.
var Platform string

// VersionText returns a message with the version.
func VersionText() string {
	return fmt.Sprintf("Commit: %s\nBuilt: %s\nGo: %s\nPlatform: %s", version.Commit, version.BuildDate, version.GoVersion, version.Platform)
}
