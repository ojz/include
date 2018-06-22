package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ojz/include/internal/README_md"
)

//go:generate include README.md

func main() {
	if len(os.Args) != 2 {
		println(string(README_md.Bytes()))
		return
	}

	filename := os.Args[1]

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	cleanFilename := strings.Replace(filename, ".", "_", -1)
	dir := "internal/" + cleanFilename
	err = os.MkdirAll(dir, os.ModeDir|0755)
	if err != nil {
		log.Fatal(err)
	}

	encoded := base64.StdEncoding.EncodeToString(content)
	source := fmt.Sprintf(tpl, filepath.Base(cleanFilename), encoded)

	err = ioutil.WriteFile(dir+"/main.go", []byte(source), 0755)
	if err != nil {
		log.Fatal(err)
	}
}

var tpl = `package %v

import (
	"encoding/base64"
)

// Bytes returns the contents of the file.
func Bytes() []byte {
	decoded, err := base64.StdEncoding.DecodeString("%v")
	if err != nil {
		panic(err)
	}
	return decoded
}
`
