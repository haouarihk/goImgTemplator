package main

import (
	"image"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
)

type PfpPos struct {
	x, y int
	size uint
}

type LevelUpPack struct {
	templateSrc      image.Image
	templateSrcShiny image.Image
	pfpPosition      PfpPos
}

type LevelUpOptions struct {
	shiny             bool
	costumeBackground image.Image
}

type ThemePack struct {
	// colors in hex
	PrimaryColorHex   string
	SecondaryColorHex string

	// default font path
	defaultFontPath string

	// fornt sizes
	bigFontSize    float64
	mediumFontSize float64
	smallFontSize  float64

	// packs
	levelUpPack LevelUpPack
}

type Templater struct {
	scale      int
	themePacks map[string]ThemePack
}

func (tp *Templater) Scale(scale int) {
	tp.scale = scale
}

// do this scaling on setup
// tepSrc = resize.Resize(uint(tepSrc.Bounds().Dx()), uint(tepSrc.Bounds().Dy()), tepSrc, resize.Lanczos2)

func (tp *Templater) levelUp(user UserInput, options LevelUpOptions, theme string) string {
	thatPack := tp.themePacks[theme].levelUpPack

	tepSrc := thatPack.templateSrc

	// if shiny
	if options.shiny {
		tepSrc = thatPack.templateSrcShiny
	}

	// get dimensions
	tepBounds := tepSrc.Bounds()

	// creating the canvas
	dc := gg.NewContext(tepBounds.Dx()*scale, tepBounds.Dy()*scale)
	dc.Clear()

	dc.Scale(float64(tp.scale), float64(tp.scale))

	// draw background
	if options.costumeBackground != nil {
		dc.DrawImage(resize.Resize(uint(tepBounds.Dx()), uint(tepBounds.Dy()), options.costumeBackground, resize.Bicubic), 0, 0)
	}

	// draw the template
	dc.DrawImage(tepSrc, 0, 0)

	// draw the pfp
	drawPfp(dc, pfpSrc, 129, 162, 351, circleSrc)

}

func LevelUpTemplating(input LevelUpInput) image.Image {

	// draw the pfp
	drawPfp(dc, pfpSrc, 129, 162, 351, circleSrc)

	// setting font
	fontSize := float64(52)
	useFont(dc, input.costumeFontPath, fontSize)

	createTextBox(dc, "CONGRATULATIONS !", secondaryColorHex, 690, 90)

	createTextBox(dc, "LEVEL", secondaryColorHex, 1060, 340)

	// settomg font
	fontSize = float64(74)
	useFont(dc, input.costumeFontPath, fontSize)

	createCenteredTextBox(dc, "USERNAME#0001", primaryColorHex, 950, 170, fontSize)

	createTextBox(dc, "YOU LEVELED UP !", "#212129", 630, 580)

	fontSize = float64(102)
	useFont(dc, input.costumeFontPath, fontSize)

	createCenteredTextBox(dc, Formatnum(input.Level), primaryColorHex, 1130, 400, fontSize)

	// save
	dc.Clip()

	return dc.Image()
}

func drawPfp(dc *gg.Context, pfpSrc image.Image, x, y, w, h int) {
	dc.DrawImage(resize.Resize(uint(w), uint(h), pfpSrc, resize.Lanczos2), x, y)
}
