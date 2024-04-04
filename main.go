package main

import (
	"log"
	"space-memes/utils"
)

func main() {

	a := &utils.ApplicationConfiguration{
		ConfigFile: "config.yaml",
	}

	a.ConfigureApplication(a.ConfigFile)

	metaFile := a.GetMetaFile()

	log.Println(metaFile)

	/*
		// eventaully want to switch this to log level based output
		if a.Debug {
			spew.Dump(a)
		}

		a.BuildFileList()

		for idx, file := range a.FileList {

			f, _ := os.Open(file)
			contentType, _ := utils.GetFileContentType(f)

			log.Printf("Photo[%v]:\t %v\t FileType: %v\n",
				idx, a.FileList[idx],
				contentType)
		}
	*/
}
