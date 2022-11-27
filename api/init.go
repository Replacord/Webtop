package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"

	"webtop/lorca"
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
	Window windoD `json:"window"`
	Web    weD    `json:"web"`
}

func ReadConfig(configFile string) *confiD {

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

func Bind(ui lorca.UI) {
	defer ui.Close()

	ui.Bind("Archive_ZipFile", Rp_zip)
	ui.Bind("Archive_UnzipFile", Rp_unzipFile)
	ui.Bind("Archive_UnzipFileRecursive", Rp_unzipSource)

	ui.Bind("Env_ClearVars", Rp_clearENV)
	ui.Bind("Env_GetVar", Rp_getENV)
	ui.Bind("Env_AddVar", Rp_addENV)
	ui.Bind("Env_RemoveVar", Rp_removeENV)

	ui.Bind("FileSys_CreateFile", Rp_createFile)
	ui.Bind("FileSys_WriteToFile", Rp_writeToFile)
	ui.Bind("FileSys_WriteFile", Rp_writeFile)
	ui.Bind("FileSys_ReadFile", Rp_readFile)
	ui.Bind("FileSys_RemoveDirectory", Rp_removeDir)
	ui.Bind("FileSys_RemoveFile", Rp_removeFile)
	ui.Bind("FileSys_FileExist", Rp_fileExist)
}
