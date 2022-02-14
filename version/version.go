package version

import "fmt"

const Version = "0.1.0"

// getVersion Compulsory minimum version, Minimum downward compatibility to this version
func getVersion() string {
	return "0.1.0"
}

// PrintVersion print currently version info
func PrintVersion() {
	fmt.Printf("Version: %s\nCore version: %s\nSame core version of generate-tool\n", Version, getVersion())
}
