package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/atotto/clipboard"
)

var semo = make(chan struct{}, 1)
var strClipboard string
var Password = "clipshoot"

func readTarget() ([]string, error) {
	rp := strings.NewReplacer(
		"'", "＇",
		"\"", "＂",
		"“", "＂",
		"”", "＂",
		",", "，",
		".", "．",
		";", "；",
		":", "：",
		"?", "？",
		"!", "！",
		"(", "（",
		")", "）",
		"[", "【",
		"]", "】",
		"{", "『",
		"}", "』",
		"+", "＋",
		"-", "－",
		"*", "＊",
		"\\", "＼",
		"/", "／",
		"<", "《",
		">", "》",
		" ", "",
		"　", "",
		"%", "％",
		"#", "＃",
		"$", "＄",
		"&", "＆",
		"=", "＝",
		"_", "＿",
		"^", "＾",
	)
	x := rp.Replace(strClipboard)
	re := regexp.MustCompile(`.*?(` + x + `).*?\|(.*)`)
	f, err := os.ReadFile(filepath.Join(".", "target.txt"))
	if err != nil {
		return nil, err
	}
	ls := re.FindAllStringSubmatch(string(f), -1)
	for _, v1 := range ls {
		for i, v2 := range v1 {
			if i == 2 {
				fmt.Printf("%s\n", v2)
			}
		}
	}
	return nil, nil
}

func verify() bool {
	text, _ := clipboard.ReadAll()
	return text == Password
}

func action() error {
	clipboard.WriteAll("") // init clipboard
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
		fmt.Printf("\n\nClipboard: %s\n", strClipboard)
	}
	return nil
}

func main() {
	if !verify() {
		return
	}
	if err := action(); err != nil {
		log.Fatal(err)
	}
}
