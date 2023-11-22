package version

import (
	"fmt"
	"time"
)

var (
	// Version of the app
	Version = "0.0.1"

	// CommitHash is the commit this version was built on, needs to be set by the linker
	CommitHash = "n/a"

	// CompileDate is the date this binary was compiled on
	CompileDate = ""
)

// BuildVersion combines available information to a nicer looking version string
func BuildVersion() string {
	var date = CompileDate
	//goland:noinspection GoBoolExpressions
	if len(date) == 0 {
		date = time.Now().String()
	}
	return fmt.Sprintf("%s-%s (%s)", Version, CommitHash, date)
}

// Ref: https://github.com/benweidig/tortuga/blob/master/version/version.go
