package gtool

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//get file relation path

//current path
func PathCurrent() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //return  filepath.Dir(os.Args[0])
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //replace \ to /
}

//current son dir
func PathCurrentSon(dirName string) string {
	return fmt.Sprintf("%s/%s", PathCurrent(), dirName)
}
