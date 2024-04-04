// General Collection of functions and utilities that will support other pacakges
package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v3"
)

// ApplicationConfig in the main configuration structure for program run time
type ApplicationConfiguration struct {
	ConfigFile string
	PictureDir string `yaml:"pictureDir"`
	IndexFile  string `yaml:"indexFile"`
	Debug      bool   `yaml:"debug"`
	FileList   []string
	Scrapper   struct {
		BaseURL   string `yaml:"baseURL"`
		IndexSlug string `yaml:"indexSlug"`
	}
}

type indexEntry struct {
	PicPath        string
	PicDescription string
}

type IndexFileConfig struct {
	FullPath   string
	IndexFile  string
	PictureDir string
	ItemsList  struct {
		Item []indexEntry
	}
}

func (a *ApplicationConfiguration) GetMetaFile() string {
	// Initialize Metadata file structure.

	idxFile := &IndexFileConfig{
		FullPath:   a.PictureDir + "/" + a.IndexFile,
		IndexFile:  a.IndexFile,
		PictureDir: a.PictureDir,
	}
	// file exist?
	file, err := os.Stat(idxFile.FullPath)
	if os.IsNotExist(err) {
		log.Println("Index File Doesn't exist Creating it now.", idxFile.FullPath)
		// Create File if it doesn't exist
		file, err := os.Create(idxFile.FullPath)
		// initalize file
		jdx, _ := json.Marshal(idxFile)
		file.WriteString(string(jdx))
		if err != nil {
			log.Println(err)
		}
		// Check format?
		if a.Debug {
			spew.Dump(file)
		}
	}
	// File Found
	if a.Debug {
		log.Println("Index File Found: ", idxFile.FullPath)
		spew.Dump(file)
	}
	return idxFile.FullPath
}

func (c *ApplicationConfiguration) ConfigureApplication(fileID string) *ApplicationConfiguration {
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

// Creates list of files within directory
func (a *ApplicationConfiguration) BuildFileList() {

	fileName := a.GetMetaFile()

	log.Println(fileName)

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
		}
	}

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
