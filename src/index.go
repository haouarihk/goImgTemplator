package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"strconv"

	"github.com/nfnt/resize"
	"golang.org/x/image/webp"

	"github.com/fogleman/gg"
)

type LevelUpInput struct {
	UserName          string
	Pfp               *os.File
	costumeBackground *os.File
	costumeFontPath   string
	Level             int
	shiny             bool
}

type LevelyInput struct {
	UserName          string
	MemberSince       string
	Pfp               *os.File
	costumeBackground *os.File
	costumeFontPath   string
	Multiplier        int
	Messages          int
	VoicesMinutes     int
	XP                int
	MaxXP             int
	TextXP            int
	VoiceXP           int
	Level             int
	shiny             bool
}

type UserInput struct {
	UserName string
	Pfp      *os.File
	Level    int
	XP       int
}

type LeaderBoardInput struct {
	FirstPlace        UserInput
	SecondPlace       UserInput
	ThirdPlace        UserInput
	costumeBackground *os.File
	costumeFontPath   string
	shiny             bool
}

const primaryColorHex string = "#04A8C3"
const secondaryColorHex string = "#bbbbbb"

var primaryColorRGBA color.Color = color.RGBA{R: 4, G: 168, B: 195, A: 255}

const scale int = 2

func main() {
	fmt.Println("Hello, World!")

	// read UIMG user image
	UIMG, _ := os.Open("UIMG.png")
	UIMG2, _ := os.Open("UIMG.png")
	UIMG3, _ := os.Open("UIMG.png")
	defer UIMG.Close()
	defer UIMG2.Close()
	defer UIMG3.Close()

	// img := LevelUpTemplating(LevelUpInput{
	// 	UserName:        "TIMG",
	// 	Pfp:             UIMG,
	// 	costumeFontPath: "fonts/SourceSansPro-Bold.ttf",
	// 	Level:           1,
	// 	shiny:           false,
	// })

	k, _ := os.Open("k.jpg")
	img := LevelyTemplating(LevelyInput{
		UserName:          "MyUSERNAME#0001",
		MemberSince:       "1/1/19",
		Pfp:               UIMG,
		costumeFontPath:   "fonts/SourceSansPro-Bold.ttf",
		costumeBackground: k,
		Multiplier:        420,
		Messages:          600,
		VoicesMinutes:     400,
		XP:                500,
		MaxXP:             14000,
		TextXP:            999999999,
		VoiceXP:           999999999,
		Level:             9009,
	})

	// img := LeaderBoardTemplating(LeaderBoardInput{
	// 	FirstPlace: UserInput{
	// 		UserName: "FIRSTPLACE#0001",
	// 		Pfp:      UIMG,
	// 		Level:    1000,
	// 		XP:       14000,
	// 	},

	// 	SecondPlace: UserInput{
	// 		UserName: "SECONDPLACE#0002",
	// 		Pfp:      UIMG2,
	// 		Level:    1,
	// 		XP:       14000,
	// 	},

	// 	ThirdPlace: UserInput{
	// 		UserName: "THRIDPLACE#0003",
	// 		Pfp:      UIMG3,
	// 		Level:    1,
	// 		XP:       1,
	// 	},

	// 	shiny: false,

	// 	costumeFontPath: "fonts/SourceSansPro-Bold.ttf",
	// })

	// save img in png format
	f, _ := os.Create("output.png")
	defer f.Close()
	jpeg.Encode(f, img, nil)
}

func LevelUpTemplating(input LevelUpInput) image.Image {
	levelUpTMPL, err := os.Open("templates/levelUp.png")

	circleTMPL, _ := os.Open("templates/circle.png")

	if input.shiny {
		levelUpTMPL, err = os.Open("templates/levelUpShiny.png")
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	defer levelUpTMPL.Close()
	defer circleTMPL.Close()

	tepSrc, _, _ := image.Decode(levelUpTMPL)

	circleSrc, _, _ := image.Decode(circleTMPL)

	// read pfp
	pfpSrc, _ := webp.Decode(input.Pfp)

	tepBounds := tepSrc.Bounds()

	// draw srcCircle to tepSrc
	// dist := image.NewRGBA(tepBounds)

	dc := gg.NewContext(tepBounds.Dx(), tepBounds.Dy())
	dc.Clear()

	// for testing
	// dc.DrawImage(image.NewUniform(color.RGBA{255, 255, 255, 255}), 0, 0)

	// draw background
	if input.costumeBackground != nil {
		costumeBackgroundSrc, _, _ := image.Decode(input.costumeBackground)
		dc.DrawImage(costumeBackgroundSrc, 0, 0)
	}

	// draw the template
	dc.DrawImage(tepSrc, 0, 0)

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

func LevelyTemplating(myInput LevelyInput) image.Image {

	// stock template image
	levelyTMPL, err := os.Open("templates/levely.png")
	// levelyBCK, err := os.Open("templates/levelyback.svg")

	circleTMPL, _ := os.Open("templates/circle.png")

	if myInput.shiny {
		// stock template image with shine
		levelyTMPL, err = os.Open("templates/levelyShiny.png")
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	defer levelyTMPL.Close()
	// defer levelyBCK.Close()
	defer circleTMPL.Close()

	tepSrc, _ := png.Decode(levelyTMPL)
	circleSrc, _ := png.Decode(circleTMPL)
	// tepbckSrc, _, _ := svg.Decode(levelyBCK)

	// scaling
	tepSrc = resize.Resize(uint(tepSrc.Bounds().Dx()), uint(tepSrc.Bounds().Dy()), tepSrc, resize.Lanczos2)

	// read pfp
	pfpSrc, _, _ := image.Decode(myInput.Pfp)

	tepBounds := tepSrc.Bounds()

	dc := gg.NewContext(tepBounds.Dx()*scale, tepBounds.Dy()*scale)
	dc.Clear()

	dc.Scale(float64(scale), float64(scale))

	// for testing
	// dc.DrawImage(image.NewUniform(color.RGBA{255, 255, 255, 255}), 0, 0)

	// draw background
	if myInput.costumeBackground != nil {
		costumeBackgroundSrc, _, _ := image.Decode(myInput.costumeBackground)
		costumeBackgroundSrc = resize.Resize(uint(tepBounds.Dx()), uint(tepBounds.Dy()), costumeBackgroundSrc, resize.Bicubic)
		dc.DrawImage(costumeBackgroundSrc, 0, 0)
	}

	// draw the template
	dc.DrawImage(tepSrc, 0, 0)
	// dc.DrawImage(tepbckSrc, 0, 0)
	dc.SetHexColor(primaryColorHex)

	dc.SetFillStyle(gg.NewSolidPattern(primaryColorRGBA))
	dc.DrawRoundedRectangle(100, 53, 690, 185, 37)

	dc.Fill()

	dc.SetFillStyle(gg.NewSolidPattern(color.Black))
	dc.DrawRoundedRectangle(200, 165, 577, 60, 25)

	dc.Fill()

	// draw the pfp
	drawPfp(dc, pfpSrc, 16, 15, 260, circleSrc)

	fontSize := float64(38)

	if err := dc.LoadFontFace(myInput.costumeFontPath, fontSize); err != nil {
		panic(err)
	}
	createTextBox(dc, "MEMBER SINCE", secondaryColorHex, 300, 208)
	createTextBox(dc, "LEVEL", secondaryColorHex, 570-(fontSize*float64(len(Formatnum(myInput.Level)))/4), 328)

	createTextBox(dc, "MULTIPLIER", "#000000", 1120, 380)
	createTextBox(dc, "MESSAGES", "#000000", 1120, 480)
	createTextBox(dc, "VOICE", "#000000", 1120, 570)
	createTextBox(dc, "MINUTES", "#000000", 1120, 600)
	fontSize = float64(48)

	if err := dc.LoadFontFace(myInput.costumeFontPath, fontSize); err != nil {
		panic(err)
	}

	createCenteredTextBox(dc, myInput.UserName, "#000000", 535, 110, fontSize)

	createTextBox(dc, "TEXT XP", secondaryColorHex, 1010, 105)
	createTextBox(dc, "VOICE XP", secondaryColorHex, 1010, 222)

	createCenteredTextBox(dc, strconv.Itoa(myInput.Multiplier)+"X", secondaryColorHex, 1045, 368, fontSize)
	createCenteredTextBox(dc, Formatnum(myInput.Messages), secondaryColorHex, 1045, 468, fontSize)
	createCenteredTextBox(dc, Formatnum(myInput.VoicesMinutes), secondaryColorHex, 1045, 568, fontSize)

	fontSize = float64(38)
	if err := dc.LoadFontFace(myInput.costumeFontPath, fontSize); err != nil {
		panic(err)
	}

	createCenteredTextBox(dc, Formatnum(myInput.TextXP), primaryColorHex, 900, 85, fontSize)

	createCenteredTextBox(dc, Formatnum(myInput.VoiceXP), primaryColorHex, 900, 205, fontSize)

	fontSize = float64(52)
	if err := dc.LoadFontFace(myInput.costumeFontPath, fontSize); err != nil {
		panic(err)
	}

	createCenteredTextBox(dc, Formatnum(myInput.Level), primaryColorHex, 730, 305, fontSize)

	levelsFontSize := float64(42)
	if err := dc.LoadFontFace(myInput.costumeFontPath, levelsFontSize); err != nil {
		panic(err)
	}

	createTextBox(dc, Formatnum(myInput.XP), primaryColorHex, 50, 330)

	maxXPstr := "/ " + Formatnum(myInput.MaxXP)
	a := float64(len(Formatnum(myInput.XP)))
	b := float64(len(maxXPstr))
	createTextBox(dc, maxXPstr, secondaryColorHex, (1.3+a/2)*levelsFontSize, 330)

	fontSize = float64(48)
	if err := dc.LoadFontFace(myInput.costumeFontPath, fontSize); err != nil {
		panic(err)
	}

	createTextBox(dc, myInput.MemberSince, primaryColorHex, 600, 211)

	createTextBox(dc, "XP", primaryColorHex, (1.2+(a+b)/2)*levelsFontSize, 330)

	drawLevelBar(dc, myInput.XP, 62, 343, float64(myInput.MaxXP), 729, 38)

	// save
	dc.Clip()
	return dc.Image()
}

func LeaderBoardTemplating(input LeaderBoardInput) image.Image {
	leaderBoardTMPL, err := os.Open("templates/leaderboard.png")

	circleTMPL, _ := os.Open("templates/circle.png")

	if input.shiny {
		leaderBoardTMPL, err = os.Open("templates/leaderboardShiny.png")
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	defer leaderBoardTMPL.Close()
	defer circleTMPL.Close()

	tepSrc, _, _ := image.Decode(leaderBoardTMPL)

	circleSrc, _, _ := image.Decode(circleTMPL)

	// read pfp

	// resizing image before making the circle so that it doesn't have that big aliasing

	const pfpSizePlace uint = 351

	tepBounds := tepSrc.Bounds()

	// draw srcCircle to tepSrc
	// dist := image.NewRGBA(tepBounds)

	dc := gg.NewContext(tepBounds.Dx(), tepBounds.Dy())
	dc.Clear()

	// for testing
	// dc.DrawImage(image.NewUniform(color.RGBA{255, 255, 255, 255}), 0, 0)

	// draw background
	if input.costumeBackground != nil {
		costumeBackgroundSrc, _, _ := image.Decode(input.costumeBackground)
		dc.DrawImage(costumeBackgroundSrc, 0, 0)
	}

	// draw the template
	dc.DrawImage(tepSrc, 0, 0)

	// draw the pfp4547745342354232423414

	fontSize := float64(52)
	useFont(dc, input.costumeFontPath, fontSize)

	createTextBox(dc, "MONTHLY", secondaryColorHex, 1000, 75)

	fontSize = float64(32)
	useFont(dc, input.costumeFontPath, fontSize)

	// first place
	drawLeaderBoardProfile(dc, input.FirstPlace, 549, 38, 253, fontSize, circleSrc)

	// second place
	drawLeaderBoardProfile(dc, input.SecondPlace, 146, 157, 253, fontSize, nil)

	// third place
	drawLeaderBoardProfile(dc, input.ThirdPlace, 947, 227, 253, fontSize, nil)

	// save
	dc.Clip()

	return dc.Image()
}

func drawLevelBar(dc *gg.Context, val int, x, y, max, maxX, maxY float64) {
	dc.SetHexColor("#ffffff")
	dc.SetFillStyle(gg.NewSolidPattern(primaryColorRGBA))

	pers := float64(val) / max
	if pers > 1 {
		pers = 1
	}

	fmt.Println(pers)

	leni := (maxX - x) * pers
	maxiY := maxY / 2
	if pers > 0 {
		dc.DrawCircle(x-2, y+maxiY, maxiY)

		dc.DrawRectangle(x-7, y, leni+14, maxY)

		dc.DrawCircle(x+10+leni, y+maxiY, maxiY)

		dc.Fill()
	}
}

func drawLeaderBoardProfile(dc *gg.Context, user UserInput, x, y int, size uint, fontSize float64, cirlceSrc image.Image) {
	pfp1Src, _ := webp.Decode(user.Pfp)
	drawPfp(dc, pfp1Src, x, y, size, cirlceSrc)

	gap := 30

	createCenteredTextBox(dc, user.UserName, primaryColorHex, x+136, y+int(size)+gap, fontSize)

	k := float64(y + int(size) + gap + int(fontSize) + 20)

	formnumLevel := Formatnum(user.Level)
	k2 := len(formnumLevel) * int(fontSize) / 2

	createTextBox(dc, formnumLevel, primaryColorHex, float64(x+100-k2), k)

	createTextBox(dc, "LVL", secondaryColorHex, float64(x+40-k2), k)

	createTextBox(dc, Formatnum(user.XP), primaryColorHex, float64(x+188), k)

	createTextBox(dc, "XP", secondaryColorHex, float64(x+140), k)
}

func drawPfp(dc *gg.Context, pfpSrc image.Image, x, y int, size uint, circleSrc image.Image) {
	pfpSrc = Circle(resize.Resize(size, size, pfpSrc, resize.Lanczos3))
	dc.DrawImage(pfpSrc, x, y)

	if circleSrc != nil {
		// circleSrc = resize.Resize(size+4, size+4, circleSrc, resize.Lanczos3)
		// dc.DrawImage(circleSrc, x-2, y-2)

		// draw a circle with a stroke
		dc.SetHexColor(primaryColorHex)
		dc.SetLineWidth(12)
		dc.SetLineCap(gg.LineCapRound)
		dc.SetLineJoin(gg.LineJoinRound)
		dc.DrawCircle(float64(x+20+int(size))/2, float64(y+15+int(size))/2, float64(int(size))/2)

		dc.Stroke()

	}
}

/*Loads the given font and size to dc, to be used onward*/
func useFont(dc *gg.Context, fontPath string, fontSize float64) {
	if err := dc.LoadFontFace(fontPath, fontSize); err != nil {
		panic(err)
	}
}

func createCenteredTextBox(dc *gg.Context, text string, textColorHex string, left int, top int, fontSize float64) *gg.Context {
	dc.SetHexColor(textColorHex)
	dc.DrawStringAnchored(text, float64(left)-(float64(len(text)/2)), float64(top), 0.5, 0.5)
	return dc
}

/* ---> Creates the TextBox, image that contains the text <---
reason: doing it this way will make it easier to control the text position
*/
func createTextBox(dc *gg.Context, text string, textColorHex string, left float64, top float64) *gg.Context {
	dc.SetHexColor(textColorHex)
	dc.DrawString(text, left, top)
	return dc
}

/** converts image.Image to []byte (PNG) **/
func pngEncodeBytes(k image.Image) []byte {
	buf := new(bytes.Buffer)
	err := png.Encode(buf, k)
	if err != nil {
		fmt.Println("unable to encode image.", err)
	}
	return buf.Bytes()
}

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
