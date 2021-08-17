package version

import (
	"fmt"
	"os"
)

var Version = ""

func PrintVersion() {
	fmt.Printf("Version: %s\n", Version)
	os.Exit(0)
}