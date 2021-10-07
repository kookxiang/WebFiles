package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	addr     = flag.String("addr", "", "HTTP server address")
	port     = flag.Int("port", 8080, "HTTP server port")
	dir      = flag.String("dir", ".", "working directory")
	basePath = flag.String("path", "/", "web access")

	baseDir string
)

type fileInfo struct {
	Name        string `json:"name"`
	Size        int64  `json:"size"`
	ModifyTime  int64  `json:"mtime"`
	IsDirectory bool   `json:"directory"`
}

func main() {
	flag.Parse()
	if filepath.IsAbs(*dir) {
		baseDir = filepath.Clean(*dir)
	} else {
		var err error
		baseDir, err = filepath.Abs(*dir)
		if err != nil {
			panic(err)
		}
	}
	if !strings.HasSuffix(*basePath, "/") {
		*basePath += "/"
	}

	serverAddr := *addr + ":" + strconv.Itoa(*port)

	fmt.Println("Start HTTP Server at", serverAddr)
	fmt.Println("Root:", baseDir)

	http.HandleFunc("/", handleRequest)

	if err := http.ListenAndServe(serverAddr, nil); err != nil {
		panic(err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path
	filePath = strings.ReplaceAll(filePath, "\\", "/")
	filePath = path.Clean(filePath)
	if !strings.HasPrefix(filePath, *basePath) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	filePath = strings.Replace(filePath, *basePath, "", 1)
	targetFile := filepath.FromSlash(filepath.Clean(filepath.Join(baseDir, filePath)))
	if !strings.HasPrefix(targetFile, baseDir) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if stat, err := os.Stat(targetFile); err != nil {
		if filePath == "App.js" {
			if _, err = os.Stat("assets/public/app.js"); err == nil {
				http.ServeFile(w, r, "assets/public/app.js")
			} else {
				w.WriteHeader(200)
				w.Header().Set("Content-Type", "application/javascript")
				_, _ = w.Write(assetsDirectoryIndexJs)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
		return
	} else if stat.IsDir() {
		if !strings.HasSuffix(r.URL.Path, "/") {
			w.Header().Set("Location", r.URL.Path+"/")
			w.WriteHeader(http.StatusTemporaryRedirect)
			return
		} else if strings.Contains(r.Header.Get("Accept"), "json") {
			data := generateDirectoryMetadata(targetFile)
			if jsonData, marshalErr := json.Marshal(data); marshalErr != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Header().Set("Content-Type", "text/plain")
				_, _ = w.Write([]byte(marshalErr.Error()))
				return
			} else {
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write(jsonData)
			}
		} else {
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "text/html")
			_, _ = w.Write([]byte(renderDirectoryIndex(*basePath)))
			return
		}
	} else {
		http.ServeFile(w, r, targetFile)
		return
	}
}

func generateDirectoryMetadata(dir string) (result []fileInfo) {
	result = make([]fileInfo, 0)
	items, _ := ioutil.ReadDir(dir)
	for _, item := range items {
		result = append(result, fileInfo{
			Name:        item.Name(),
			Size:        item.Size(),
			IsDirectory: item.IsDir(),
			ModifyTime:  item.ModTime().UnixMilli(),
		})
	}
	return
}
