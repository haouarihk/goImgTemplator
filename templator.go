package Templator

import (
	"image"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/nfnt/resize"
)

type ImageObject struct {
	Src          image.Image
	X, Y         int
	W, H         uint
	Centered     bool
	RightAligned bool
}

type TextObject struct {
	Text             string
	Color            string
	X, Y             float64
	FontFace         *truetype.Font
	FontSize         float64
	hiddenAddedText  string
	hiddenSubbedText string
	Centered         bool
	RightAligned     bool
	TextAfter        *TextObject
	TextBefore       *TextObject
}

type UserTemplate struct {
	Pfp          ImageObject
	Username     TextObject
	Level        TextObject
	XP           TextObject
	XPAndMaxXP   TextObject
	XpBar        XpBar
	MaxXp        TextObject
	MemberSince  TextObject
	TextXP       TextObject
	VoiceXP      TextObject
	Multiplier   TextObject
	Messages     TextObject
	VoiceMinutes TextObject
}

type UserInput struct {
	Username     string
	Tag          string
	Pfp          image.Image
	Level        int
	XP           int
	MaxXP        int
	MemberSince  string
	TextXP       int
	VoiceXP      int
	Multiplier   int
	Messages     int
	VoiceMinutes int
}

type Options struct {
	shiny             bool
	costumeBackground image.Image
	constumeFontFace  *truetype.Font
	theme             string
	pack              string
}

type ThemePack struct {
	TemplateSrc      image.Image
	TemplateSrcShiny image.Image
	Scale            int
	Texts            []TextObject
	images           []ImageObject
	UserTemplate     []UserTemplate
}

type Theme struct {
	// colors in hex
	PrimaryColorHex   string
	SecondaryColorHex string

	// default font path
	DefaultFontFace *truetype.Font
	DefaultFontSize float64

	// fornt sizes
	bigFontSize    float64
	mediumFontSize float64
	smallFontSize  float64

	// Packs
	Packs map[string]ThemePack
}

type Templator struct {
	themes map[string]Theme
}

type XpBar struct {
	X          float64
	Y          float64
	Width      float64
	Height     float64
	Roundyness float64
	Color      string
	hideText   bool

	// optional default values
	XP    int
	MaxXP int
}

func (tp *Templator) render(user []UserInput, options Options) image.Image {
	thatPack := tp.themes[options.theme].Packs[options.pack]

	tepSrc := thatPack.TemplateSrc

	// if shiny
	if options.shiny {
		tepSrc = thatPack.TemplateSrcShiny
	}

	// default font
	defaultFontFace := tp.themes[options.theme].DefaultFontFace

	// if options.constumeFontFace != nil {
	// 	defaultFontFace = options.constumeFontFace
	// }

	// get dimensions
	tepBounds := tepSrc.Bounds()
	scale := thatPack.Scale | 1
	// creating the canvas
	dc := gg.NewContext(tepBounds.Dx()*scale, tepBounds.Dy()*scale)
	// dc.Clear()

	// scaling the canvas
	dc.Scale(float64(scale), float64(scale))

	// draw background
	if options.costumeBackground != nil {
		dc.DrawImage(resize.Resize(uint(tepBounds.Dx()), uint(tepBounds.Dy()), options.costumeBackground, resize.Bicubic), 0, 0)
	}

	// draw the template
	dc.DrawImage(tepSrc, 0, 0)

	// drawing static texts
	for _, text := range thatPack.Texts {
		drawText(dc, &text, defaultFontFace)
	}

	// ---- draw user properties
	for i, userTemplate := range thatPack.UserTemplate {
		emptyTextBox := TextObject{}
		useri := user[i]

		// draw username
		if userTemplate.Username != emptyTextBox {
			_username := userTemplate.Username
			_username.Text = useri.Username

			if userTemplate.Username != emptyTextBox {
				_username.TextAfter.Text = useri.Tag
			}

			drawText(dc, &_username, defaultFontFace)
		}

		// draw the pfp
		if useri.Pfp != nil {
			userTemplate.Pfp.Src = Circle(useri.Pfp)
			drawImage(dc, &userTemplate.Pfp)
		}

		// draw XP Bar
		if (userTemplate.XpBar != XpBar{}) {
			_xpbar := userTemplate.XpBar
			_xpbar.XP = useri.XP
			_xpbar.MaxXP = useri.MaxXP
			drawXpBar(dc, _xpbar)
		}

		// draw XPMaxXP
		if userTemplate.XPAndMaxXP != emptyTextBox {
			_XPAndMaxXP := userTemplate.XPAndMaxXP
			_XPAndMaxXP.Text = tightyNumbers(useri.XP)

			_XPAndMaxXP.TextAfter.Text += tightyNumbers(useri.MaxXP)

			drawText(dc, &_XPAndMaxXP, defaultFontFace)
		}

		// draw XP
		if userTemplate.XP != emptyTextBox {
			_XP := userTemplate.XP
			_XP.Text = tightyNumbers(useri.XP)

			drawText(dc, &_XP, defaultFontFace)
		}

		// draw level
		if userTemplate.Level != emptyTextBox {
			_level := userTemplate.Level
			_level.Text = tightyNumbers(useri.Level)

			drawText(dc, &_level, defaultFontFace)
		}

		// draw MemberSince
		if userTemplate.MemberSince != emptyTextBox {
			_MemberSince := userTemplate.MemberSince
			_MemberSince.Text = useri.MemberSince

			drawText(dc, &_MemberSince, defaultFontFace)
		}

		// draw Text XP
		if userTemplate.TextXP != emptyTextBox {
			_textXP := userTemplate.TextXP
			_textXP.Text = tightyNumbers(useri.TextXP)

			drawText(dc, &_textXP, defaultFontFace)
		}

		// draw VOICE XP
		if userTemplate.VoiceXP != emptyTextBox {
			_voiceXP := userTemplate.VoiceXP
			_voiceXP.Text = tightyNumbers(useri.VoiceXP)

			drawText(dc, &_voiceXP, defaultFontFace)
		}

		// draw MULTIPLIER
		if userTemplate.Multiplier != emptyTextBox {
			_multiplier := userTemplate.Multiplier
			_multiplier.Text = tightyNumbers(useri.Multiplier)

			drawText(dc, &_multiplier, defaultFontFace)
		}

		// draw MESSAGES
		if userTemplate.Messages != emptyTextBox {
			_messages := userTemplate.Messages
			_messages.Text = tightyNumbers(useri.Messages)

			drawText(dc, &_messages, defaultFontFace)
		}

		// draw VOICEMINUTES XP
		if userTemplate.VoiceMinutes != emptyTextBox {
			_voiceMinutes := userTemplate.VoiceMinutes
			_voiceMinutes.Text = tightyNumbers(useri.VoiceMinutes)

			drawText(dc, &_voiceMinutes, defaultFontFace)
		}
	}

	return dc.Image()
}

func drawXpBar(dc *gg.Context, xpBar XpBar) {
	pers := float64(xpBar.XP) / float64(xpBar.MaxXP)
	if pers > 0 {
		if pers > 1 {
			pers = 1
		}

		if pers < 0.018 {
			pers = 0.018
		}

		leni := xpBar.Width * pers

		dc.SetHexColor(xpBar.Color)
		dc.SetFillStyle(gg.NewSolidPattern(ParseHexColorFast(xpBar.Color)))
		dc.DrawRoundedRectangle(xpBar.X, xpBar.Y, leni+14, xpBar.Height, xpBar.Roundyness)
		dc.Fill()
	}
}

func drawLevelBar(dc *gg.Context, val int, x, y, max, maxX, maxY float64) {
	dc.SetHexColor("#ffffff")
}

func drawText(dc *gg.Context, text *TextObject, defaultFont *truetype.Font) {
	DefaultingFontFaceOfTextObject(text, defaultFont)
	fontface := easyGetFontFace(text.FontFace, text.FontSize)
	dc.SetFontFace(fontface)
	w, _ := dc.MeasureString(text.Text)

	if text.Centered {
		// draw text before
		if text.TextBefore != nil {
			text.TextBefore.X += text.X - (w / 2)
			text.TextBefore.Y += text.Y
			drawText(dc, text.TextBefore, defaultFont)
		}

		dc.SetFontFace(fontface)
		dc.SetHexColor(text.Color)
		dc.DrawStringAnchored(text.Text, text.X, text.Y, 0.5, 0.5)

		// draw text after
		if text.TextAfter != nil {
			text.TextAfter.X += text.X + (w / 2)
			text.TextAfter.Y += text.Y
			drawText(dc, text.TextAfter, defaultFont)
		}
	} else if text.RightAligned {
		// draw text before
		if text.TextBefore != nil {
			text.TextBefore.X += text.X - w
			text.TextBefore.Y += text.Y
			drawText(dc, text.TextBefore, defaultFont)
		}

		dc.SetFontFace(fontface)
		dc.SetHexColor(text.Color)
		dc.DrawStringAnchored(text.Text, text.X, text.Y, 1, 0.5)

		// draw text after
		if text.TextAfter != nil {
			text.TextAfter.X += text.X
			text.TextAfter.Y += text.Y
			drawText(dc, text.TextAfter, defaultFont)
		}
	} else {
		// draw text before
		if text.TextBefore != nil {
			text.TextBefore.X += text.X
			text.TextBefore.Y += text.Y
			drawText(dc, text.TextBefore, defaultFont)
		}

		dc.SetFontFace(fontface)
		dc.SetHexColor(text.Color)
		dc.DrawStringAnchored(text.Text, text.X, text.Y, 0, 0.5)

		// draw text after
		if text.TextAfter != nil {
			text.TextAfter.X += text.X + w
			text.TextAfter.Y += text.Y
			drawText(dc, text.TextAfter, defaultFont)
		}
	}
}

func drawImage(dc *gg.Context, image *ImageObject) {
	if image.Centered {
		dc.DrawImageAnchored(resize.Resize(image.W, image.H, Circle(image.Src), resize.Lanczos2), image.X, image.Y, 0.5, 0.5)
	} else if image.RightAligned {
		dc.DrawImageAnchored(resize.Resize(image.W, image.H, Circle(image.Src), resize.Lanczos2), image.X, image.Y, 1, 0.5)
	} else {
		dc.DrawImageAnchored(resize.Resize(image.W, image.H, Circle(image.Src), resize.Lanczos2), image.X, image.Y, 0, 0)
	}
}
