package Templator

import (
	"image"
	"image/color"
	"image/draw"
	"io/ioutil"
	"strconv"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

type circle struct {
	centerPoint image.Point
	radius      int
}

func (c *circle) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *circle) Bounds() image.Rectangle {
	return image.Rect(
		c.centerPoint.X-c.radius,
		c.centerPoint.Y-c.radius,
		c.centerPoint.X+c.radius,
		c.centerPoint.Y+c.radius,
	)
}

func (c *circle) At(x, y int) color.Color {
	xpos := float64(x-c.centerPoint.X) + 0.5
	ypos := float64(y-c.centerPoint.Y) + 0.5
	radiusSquared := float64(c.radius * c.radius)
	if xpos*xpos+ypos*ypos < radiusSquared {
		return color.Alpha{255}
	}
	return color.Alpha{0}
}

func Circle(src image.Image) image.Image {
	dst := image.NewRGBA(src.Bounds())
	r := src.Bounds().Dx() / 2
	p := image.Point{
		X: src.Bounds().Dx() / 2,
		Y: src.Bounds().Dy() / 2,
	}
	draw.DrawMask(dst, dst.Bounds(), src, image.ZP, &circle{p, r}, image.ZP, draw.Over)
	return dst
}

/* Format number to be in shape of 123,123*/
func Formatnum(n int) string {
	in := strconv.Itoa(n)
	numOfDigits := len(in)
	if n < 0 {
		numOfDigits-- // First character is the - sign (not a digit)
	}
	numOfCommas := (numOfDigits - 1) / 3

	out := make([]byte, len(in)+numOfCommas)
	if n < 0 {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = ','
		}
	}
}

/* Format number to be in shape of 99,999 or 100k*/
func TightyNumbers(n int) string {
	if n < 100000 {
		return Formatnum(n)
	} else if n < 1000000 {
		return Formatnum(n/1000) + "K"
	} else if n < 1000000000 {
		return Formatnum(n/1000000) + "M"
	}
	return Formatnum(n/1000000000) + "B"
}

/* Easy way to load a font face, it doens't return an error if the font doesn't exist or doesn't work. Make sure that the font works. And if you're not sure about it
use ``go
	fontBytes, err := ioutil.ReadFile(fontPath)
	if err != nil {
		panic(err)
	}
	fontFace, err := truetype.Parse(fontBytes)
	if err != nil {
		panic(err)
	}
``
to load the font
*/
func EasyLoadFontFace(fontPath string) (font *truetype.Font) {
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

/* An Easy way to change the font size and get font.Face out of Truetype.Font*/
func EasyGetFontFace(font *truetype.Font, size float64) (fontFace font.Face) {
	return truetype.NewFace(font, &truetype.Options{
		Size: size,
		DPI:  72,
	})
}

/* Converts Color to Hex */
func ConvertToHex(s string) (c color.RGBA) {
	c.A = 0xff

	if s[0] != '#' {
		return c
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		return 0
	}

	switch len(s) {
	case 7:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	}
	return
}

/* You give it a text and it will default it to the parent's font, if it doesn't have a font (it's used in the engine) */
func DefaultingFontOfTextObject(textObject *TextObject, defaultFontFace *truetype.Font, defaultFontSize float64) {
	// if it doesn't have a font face
	if textObject.FontFace == nil {
		textObject.FontFace = defaultFontFace
	}
	if textObject.FontSize == 0 {
		textObject.FontSize = defaultFontSize
	}

	if textObject.TextBefore != nil {
		DefaultingFontOfTextObject(textObject.TextBefore, defaultFontFace, defaultFontSize)
	}
	if textObject.TextAfter != nil {
		DefaultingFontOfTextObject(textObject.TextAfter, defaultFontFace, defaultFontSize)
	}

}
