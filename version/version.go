package version

import "fmt"

import "embed"

// Embedded contains embedded scripts.
//go:embed *
var Embedded embed.FS

var Version = setVersion()

func setVersion() string {
	rb, err := Embedded.ReadFile("version")
	if err != nil {
		return "v0.0.0"
	}
	return string(rb)
}

// getVersion Compulsory minimum version, Minimum downward compatibility to this version
func getVersion() string {
	return Version
}

// PrintVersion print currently version info
func PrintVersion() {
	fmt.Printf("Version: %s\nCore version: %s\nSame core version of mss-boot-generator\n", Version, getVersion())
}
