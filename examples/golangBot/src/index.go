package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"time"

	Templator "github.com/haouarihk/goImgTemplator/src"
)

func main() {
	fmt.Println("Hello, world!")
	UIMG, _ := os.Open("imgs/UIMG.png")
	k, _ := os.Open("imgs/k.jpg")
	defer UIMG.Close()

	UIMGSrc, _, _ := image.Decode(UIMG)
	kSrc, _, _ := image.Decode(k)

	templator := Templator.Init(setupThemes())

	start := time.Now()

	img := templator.Render([]Templator.UserInput{{
		Username:     "USERNAME",
		Tag:          "#0001",
		Pfp:          UIMGSrc,
		Level:        1000000,
		XP:           60,
		MaxXP:        100,
		TextXP:       50000,
		VoiceXP:      500000,
		VoiceMinutes: 5000000,
		Messages:     50000000,
		Multiplier:   50,
		MemberSince:  "01/10/01",
	}, {
		Username:    "USERNAME",
		Tag:         "#0002",
		Pfp:         UIMGSrc,
		Level:       1,
		XP:          5000,
		MemberSince: "01/10/01",
	}, {
		Username:    "USERNAME",
		Tag:         "#0003",
		Pfp:         UIMGSrc,
		Level:       1,
		XP:          5000,
		MemberSince: "01/10/01",
	}}, Templator.Options{
		Shiny:             false,
		CostumeBackground: kSrc,
		Theme:             "default",
		Pack:              "leaderboard",
	})

	fmt.Println("execultion time:", time.Since(start))
	// save img
	f, _ := os.Create("output.png")
	defer f.Close()
	png.Encode(f, img)

}
