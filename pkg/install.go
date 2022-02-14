package pkg

import (
	"fmt"
	"github.com/lwnmengjing/micro-service-gen-tool/version"
	"os"
	"strings"
)

// Update update generate-tool
func Update() {
	destPath := downloadLatest()
	//复制文件到对应目录
	copyStaticFile(destPath, "generate-tool")
	fmt.Println("Update completed")
}

// Upgrade check update
func Upgrade() {
	if GetLatestVersion() != version.Version {
		//need update
		fmt.Printf("do you need to upgrade[%s]: ", Yellow("y/n"))
		var upgrade bool
		var input string
		_, _ = fmt.Scan(&input)
		switch strings.ToLower(input) {
		case "y", "yes", "t", "true":
			upgrade = true
		}
		if upgrade {
			Update()
			os.Exit(0)
		}
	}
}
