package main

import (
    "fmt"
    "os"
    "net/http"
    "path/filepath"
		//"reflect"
)

func main() {
    var files []string

    root := "/home/zack-cp/projects/nasa_pics/pictures"

    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        files = append(files, path)
        return nil
    })
    if err != nil {
        panic(err)
    }

    for _, file := range files {
        //fmt.Println(reflect.TypeOf(file).String()+":\t"+file)

        f, err := os.Open(file)
				//fmt.Println(reflect.TypeOf(f).String()+":\t")
      	if err != nil {
      		panic(err)
      	}
        // Get the content

      	contentType, err := GetFileContentType(f)
				// why is this being weird?
				if err != nil {
          fmt.Println(err)
      		//panic(err)
      	}

				if  (contentType != "image/jpeg" && contentType != "image/png" && contentType != "image/gif") {
      		//fmt.Println("File:\t"+file+"\tContent Type:\t" + contentType)
				}

      	f.Close()
    }
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
  fmt.Println(contentType)
	return contentType, nil
}
