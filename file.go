package gtool

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

var fileMimeMap = map[string]string{
	"image/jpeg":               ".jpeg",
	"image/pjpeg":              ".jpeg",
	"image/jpg":                ".jpg",
	"image/png":                ".png",
	"image/gif":                ".gif",
	"image/webp":               ".webp",
	"image/bmp":                ".bmp",
	"video/mp4":                ".mp4",
	"application/octet-stream": ".jpg",
}

//get file type by content type 
func FileTypeByContentType(key string) (res string) {
	fileType := fileMimeMap[key]

	if fileType != "" {
		return fileType
	}
	return ""
}

//file dir relation functions
func FileDirAutoCreate(path string) (res bool) {
	if !FileDirIsExist(path) {
		fmt.Println("FilePathAutoCreate.noExist", path)

		err := os.MkdirAll(path, 0755)

		if err != nil {
			fmt.Println("Error", "FilePathAutoCreate", err)
			return
		}

		res = true
	}

	return
}

//file dir is exist
func FileDirIsExist(path string) (res bool) {
	_, err := os.Stat(path)

	if err == nil {
		res = true
	}

	return
}

//write content to file
func FileWrite(path, content string) bool {
	b := []byte(content + "\n")
	fd, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(TimeNow(), "Error", "write.file.err", err)
		return false
	}

	fd.Write(b)
	fd.Close()

	return true
}

//file base64 encode
func FileBase64Encode(filePath string) (base64Str string) {
	buffFile, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Println("read.file.error", err)
		return
	}

	//buff := make([]byte, 500000)
	base64Str = base64.StdEncoding.EncodeToString(buffFile)

	return
}

//file base64 decode
func FileBase64Decode(code, dest string) error {
	buff, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(dest, buff, 07440); err != nil {
		return err
	}

	return nil
}
