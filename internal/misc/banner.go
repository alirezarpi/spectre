package internal

import (
	"fmt"
	"io/ioutil"
    
	"spectre/pkg/log"
)

func ShowBanner(filepath string) {
	content, err := ioutil.ReadFile("banner.txt")
    if err != nil {
		log.Fatal("", err.Error())
		return nil
	}
    fmt.Println("\n" + string(content))
}
