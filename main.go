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
	text, err := clipboard.ReadAll()
	if err != nil {
		return err
	}
	if strClipboard != text {
		strClipboard = text
		fmt.Printf("Clipboard: %s\n", strClipboard)
		semo <- struct{}{}
	}
	return nil
}

func main() {
	if err := action(); err != nil {
		log.Fatal(err)
	}
}
