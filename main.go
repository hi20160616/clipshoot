package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/atotto/clipboard"
)

var semo = make(chan struct{}, 1)
var strClipboard string

func readTarget() ([]string, error) {
	f, err := os.ReadFile(filepath.Join(".", "target.txt"))
	if err != nil {
		return nil, err
	}
	c := string(f)
	// strClipboard = strings.ReplaceAll(strClipboard, "\\", "＼")
	re := regexp.MustCompile(`.*?(` + strClipboard + `).*?\|(.*)`)
	ls := re.FindAllStringSubmatch(c, -1)
	for _, v1 := range ls {
		for i, v2 := range v1 {
			if i == 2 {
				fmt.Printf("%s\n", v2)
			}
		}
	}
	return nil, nil
}

func action() error {
	for {
		if err := listen(); err != nil {
			return err
		}
		select {
		case <-semo:
			readTarget()
		default:
		}
	}
}

func listen() error {
	// TODO: treat err
	text, _ := clipboard.ReadAll()
	if strClipboard != text {
		strClipboard = text
		semo <- struct{}{}
		fmt.Printf("Clipboard: %s\n", strClipboard)
	}
	return nil
}

func main() {
	if err := action(); err != nil {
		log.Fatal(err)
	}
}
