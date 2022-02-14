package pkg

import "fmt"

func Update() {
	destPath := downloadLatest()
	//复制文件到对应目录
	copyStaticFile(destPath, "generate-tool")
	fmt.Println("Update completed")
}
