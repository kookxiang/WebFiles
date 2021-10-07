package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed assets/public/index.html
var assetsDirectoryIndexHtml string

//go:embed assets/public/app.js
var assetsDirectoryIndexJs []byte

func renderDirectoryIndex(basePath string) string {
	var result = assetsDirectoryIndexHtml
	result = strings.Replace(result, "/App.js", basePath+"App.js", 1)
	if basePath != "/" {
		result = strings.Replace(result, "<script", "<script>window.BASE_PATH = "+strconv.Quote(basePath)+"</script>\n\t<script", 1)
	}
	return result
}
