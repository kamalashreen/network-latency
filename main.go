package main

import (
	"fmt"
	"kamalashreen/network-latency/cmd"
	"os"
	"path/filepath"
)

// Version and build information.
var (
	version = "N/A"
	commit  = "N/A"
)

func main() {
	execPath, err := os.Executable()
	if err == nil {
		_ = os.Chdir(filepath.Dir(execPath))
	}

	version := fmt.Sprintf("%s (Commit: %s)", version, commit)
	cmd.Execute(version)
}
