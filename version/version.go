package version

import "time"

var (
	BuildTime = time.Now()
	Commit    = "unset"
	Release   = "unset"
)
