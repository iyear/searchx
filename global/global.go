package global

import _ "embed"

//go:embed logo
var Logo string

// vars below are set by '-X' flag
var (
	Version    = "dev"
	Commit     = "unknown"
	CommitDate = "unknown"
)

const DocsPath = "docs"
