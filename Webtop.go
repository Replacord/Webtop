package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"

	"github.com/zserge/lorca"
	"webtop/api/archive"
	"webtop/api/env"
	"webtop/api/filesystem"
)

type windoD struct {
	Height int `json:"height"`
	Width  int `json:"width"`
}

type weD struct {
	StaticDir  string `json:"staticDir"`
	BrowserOut string `json:"BrowserOutput"`
}

type confiD struct {
	Window windoD   `json:"window"`
	Web    weD      `json:"web"`
	libs   []string `json:"libs"`
}

func readConfig(configFile string) *confiD {

	conf := &confiD{}

	jsonFile, _ := os.Open(configFile)

	defer jsonFile.Close()

	Data, _ := ioutil.ReadAll(jsonFile)

	r := bytes.NewReader(Data)
	decoder := json.NewDecoder(r)
	decoder.Decode(conf)

	json.Unmarshal(Data, &conf)

	return conf
}

func UII() lorca.UI {

	conf := readConfig("./Webtop.json")
	BrowserOutPath, _ := filepath.Abs(conf.Web.BrowserOut)
	ui, _ := lorca.New("", BrowserOutPath, conf.Window.Height, conf.Window.Width)

	return ui

}

var ui = UII()

func main() {

	conf := readConfig("./Webtop.json")
	defer ui.Close()

	ui.Bind("rp_zip", archive.Rp_zip)
	ui.Bind("rp_unzipFile", archive.Rp_unzipFile)
	ui.Bind("rp_unzipSource", archive.Rp_unzipSource)

	ui.Bind("rp_clearENV", env.Rp_clearENV)
	ui.Bind("rp_getENV", env.Rp_getENV)
	ui.Bind("rp_addENV", env.Rp_addENV)
	ui.Bind("rp_removeENV", env.Rp_removeENV)

	ui.Bind("rp_createFile", filesystem.Rp_createFile)
	ui.Bind("rp_writeToFile", filesystem.Rp_writeToFile)
	ui.Bind("rp_writeFile", filesystem.Rp_writeFile)
	ui.Bind("rp_readFile", filesystem.Rp_readFile)
	ui.Bind("rp_removeDir", filesystem.Rp_removeDir)
	ui.Bind("rp_removeFile", filesystem.Rp_removeFile)

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	go http.Serve(ln, http.FileServer(http.Dir(conf.Web.StaticDir)))
	ui.Load(fmt.Sprintf("http://%s", ln.Addr()))
	<-ui.Done()

}
