package main

import (
	"log"
	"net/http"
	"os"
	//"reflect"
)

type ApplicationConfiguration struct {
	ConfigFile string
	PictureDir string
}

func GetFileContentType(out *os.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}
	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	log.Println(contentType)

	return contentType, nil
}

func (a *ApplicationConfiguration) buildFileList() []string {
	var fileList []string
	f, err := os.Open(a.PictureDir)
	if err != nil {
		log.Println(err)

	}
	files, err := f.Readdir(0)
	if err != nil {
		log.Println(err)

	}

	for _, v := range files {

		if !v.IsDir() {
			fileList = append(fileList, v.Name())
			log.Printf("File Path: %v/%v", a.PictureDir, v.Name())
		}
	}
	return fileList
}

func main() {

	a := &ApplicationConfiguration{
		ConfigFile: "config.yaml",
		PictureDir: "/home/zack/Pictures",
	}

	log.Println(a.buildFileList())
}
