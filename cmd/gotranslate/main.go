package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/coc1961/gotranslate/internal/translate"
)

func main() {
	language := flag.String("l", "en", "language default (en)")
	sourceLanguage := flag.String("s", "es", "source language default (es)")
	inputSource := flag.String("i", "", "input source dir")
	outputSource := flag.String("o", "", "output source dir")
	flag.Parse()

	if *inputSource == "" || *outputSource == "" {
		fmt.Println("input and output source dir are mandatory")
		flag.PrintDefaults()
		os.Exit(1)
	}
	translate, err := translate.New(*language, *sourceLanguage, *inputSource, *outputSource)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = translate.Translate()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
