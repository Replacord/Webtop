package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"path/filepath"
	"webtop/api"

	"webtop/lorca"
)

func UII() lorca.UI {

	conf := api.ReadConfig("./Wt.json")
	BrowserOutPath, _ := filepath.Abs(conf.Web.BrowserOut)
	ui, _ := lorca.New("", BrowserOutPath, conf.Window.Height, conf.Window.Width)

	return ui

}

var ui = UII()

func main() {

	conf := api.ReadConfig("./Wt.json")
	defer ui.Close()

	ui.Bind("Archive_ZipFile", api.Rp_zip)
	ui.Bind("Archive_UnzipFile", api.Rp_unzipFile)
	ui.Bind("Archive_UnzipFileRecursive", api.Rp_unzipSource)

	ui.Bind("Env_ClearVars", api.Rp_clearENV)
	ui.Bind("Env_GetVar", api.Rp_getENV)
	ui.Bind("Env_AddVar", api.Rp_addENV)
	ui.Bind("Env_RemoveVar", api.Rp_removeENV)

	ui.Bind("FileSys_CreateFile", api.Rp_createFile)
	ui.Bind("FileSys_WriteToFile", api.Rp_writeToFile)
	ui.Bind("FileSys_WriteFile", api.Rp_writeFile)
	ui.Bind("FileSys_ReadFile", api.Rp_readFile)
	ui.Bind("FileSys_RemoveDirectory", api.Rp_removeDir)
	ui.Bind("FileSys_RemoveFile", api.Rp_removeFile)
	ui.Bind("FileSys_FileExist", api.Rp_fileExist)

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	go http.Serve(ln, http.FileServer(http.Dir(conf.Web.StaticDir)))
	ui.Load(fmt.Sprintf("http://%s", ln.Addr()))
	<-ui.Done()

}
