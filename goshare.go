package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	sharePrefix = "https://play.golang.org/p/"
	shareUrl    = "https://play.golang.org/share"
	contentType = "application/x-www-form-urlencoded"
)

func exit(s string) {
	os.Stderr.WriteString(s)
	os.Stderr.WriteString("\n")
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		exit("usage: goshare <filename>")
	}

	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		exit(err.Error())
	}

	req := bytes.NewReader(b)
	res, err := http.Post(shareUrl, contentType, req)
	defer res.Body.Close()
	if err != nil {
		exit(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		exit(err.Error())
	}

	fmt.Println(sharePrefix + string(body))
}
