package utils

import (
	"log"
	"net/http"
	"os"
)

type ApplicationConfiguration struct {
	ConfigFile string
	PictureDir string
	FileList   []string
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

	//log.Println(contentType)

	return contentType, nil
}

func (a *ApplicationConfiguration) BuildFileList() {
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
			a.FileList = append(a.FileList, a.PictureDir+"/"+v.Name())
			//log.Printf("File Path: %v/%v", a.PictureDir, v.Name())
		}
	}
}
