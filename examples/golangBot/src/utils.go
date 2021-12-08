package main

import (
	"io/ioutil"

	"github.com/golang/freetype/truetype"
)

func easyLoadFontFace(fontPath string) *truetype.Font {
	fontBytes, err := ioutil.ReadFile(fontPath)
	if err != nil {
		panic(err)
	}
	f, err := truetype.Parse(fontBytes)
	if err != nil {
		panic(err)
	}
	return f
}