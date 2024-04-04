// General Collection of functions and utilities that will support other pacakges
package utils

import (
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

// ApplicationConfig in the main configuration structure for program run time
type ApplicationConfiguration struct {
	ConfigFile string
	PictureDir string `yaml:"pictureDir"`
	IndexFile  string `yaml:"indexFile"`
	FileList   []string
}

func (c *ApplicationConfiguration) ReadConfig(fileID string) *ApplicationConfiguration {
	yamlFile, err := os.ReadFile(fileID)
	if err != nil {
		log.Fatal("Err in os.ReadFile")
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatal("Failure to Unmarshal")
	}
	return c
}

// Determines the mime content-type
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

// Creates list of files within directory
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
