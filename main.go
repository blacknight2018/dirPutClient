package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var server, remote, local *string

func postFile(serverApi, dstSaveDir, FileRelPath, localFileAbsPath string) string {
	url := url.Values{}
	var ret string
	bytes, _ := ioutil.ReadFile(localFileAbsPath)
	url.Set("file_path", FileRelPath)
	url.Set("data", base64.StdEncoding.EncodeToString(bytes))
	url.Set("dir", dstSaveDir)
	resp, _ := http.Post(serverApi, "application/x-www-form-urlencoded", strings.NewReader(url.Encode()))
	if resp != nil {
		bytes, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		ret = string(bytes)
	}
	return ret
}
func f(dirPath string, preDir string) {
	if dirPath[len(dirPath)-1:] != "/" {
		dirPath = dirPath + "/"
	}
	fs, err := ioutil.ReadDir(dirPath)
	if err == nil {
		for _, v := range fs {
			if v.IsDir() {
				f(dirPath+v.Name(), preDir+"/"+v.Name())
			} else {
				//filePath := dirPath + v.Name()
				fmt.Println("put ", dirPath+v.Name(), " ", postFile(*server, *remote, preDir+"/"+v.Name(), dirPath+v.Name()))
			}
		}
	}
}
func main() {
	server = flag.String("server", "", "dirPutServer http")
	remote = flag.String("remote", "", "server dir")
	local = flag.String("local", "", "local dir")
	flag.Parse()
	fmt.Println(*server, *remote, *local)
	f(*local, "")
}
