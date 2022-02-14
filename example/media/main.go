package main

import (
	"fmt"
	"mime"
	"path/filepath"
)

func main() {
	fmt.Println(mime.TypeByExtension(filepath.Ext("123")))
}
