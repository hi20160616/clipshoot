package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	// "golang.design/x/clipboard"
	"github.com/zyedidia/clipper"
)

var semo = make(chan struct{}, 1)
var strClipboard string

func readTarget() ([]string, error) {
	f, err := os.ReadFile(filepath.Join(".", "target.txt"))
	if err != nil {
		return nil, err
	}
	c := string(f)
	if strClipboard == "" {
		return nil, nil
	}
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
	// raw := clipboard.Read(clipboard.FmtText)
	// text := string(raw)
	clip, err := clipper.GetClipboard(clipper.Clipboards...)
	if err != nil {
		return err
	}
	data, err := clip.ReadAll(clipper.RegClipboard)
	if err != nil {
		return err
	}
	text := string(data)

	if strClipboard != text {
		strClipboard = text
		semo <- struct{}{}
		fmt.Printf("\n\nClipboard: %s\n", strClipboard)
	}
	return nil
}

func main() {
	// Init returns an error if the package is not ready for use.
	// err := clipboard.Init()
	// if err != nil {
	//         panic(err)
	// }

	if err := action(); err != nil {
		log.Fatal(err)
	}
}
