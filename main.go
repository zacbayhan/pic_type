package main

import (
	"log"
	"os"

	//"reflect"
	"space-memes/utils"
)

func main() {

	a := &utils.ApplicationConfiguration{
		ConfigFile: "config.yaml",
	}
	a.ReadConfig(a.ConfigFile)

	log.Println(a)

	a.BuildFileList()

	for idx, file := range a.FileList {

		f, _ := os.Open(file)
		contentType, _ := utils.GetFileContentType(f)

		log.Printf("Photo[%v]:\t %v\t FileType: %v\n", idx, a.FileList[idx], contentType)
	}
}
