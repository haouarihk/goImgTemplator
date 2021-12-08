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
	Shiny             bool
	CostumeBackground image.Image
	ConstumeFontFace  *truetype.Font
	Theme             string
	Pack              string
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

type Templater struct {
	Themes map[string]Theme
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

func (tp *Templater) Render(user []UserInput, options Options) image.Image {
	thatPack := tp.Themes[options.Theme].Packs[options.Pack]

	tepSrc := thatPack.TemplateSrc

	// if shiny
	if options.Shiny {
		tepSrc = thatPack.TemplateSrcShiny
	}

	// default font
	defaultFontFace := tp.Themes[options.Theme].DefaultFontFace
	defaultFontSize := tp.Themes[options.Theme].DefaultFontSize
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
	if options.CostumeBackground != nil {
		dc.DrawImage(resize.Resize(uint(tepBounds.Dx()), uint(tepBounds.Dy()), options.CostumeBackground, resize.Bicubic), 0, 0)
	}

	// draw the template
	dc.DrawImage(tepSrc, 0, 0)

	// drawing static texts
	for _, text := range thatPack.Texts {
		DrawText(dc, &text, defaultFontFace, defaultFontSize)
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

			DrawText(dc, &_username, defaultFontFace, defaultFontSize)
		}

		// draw the pfp
		if useri.Pfp != nil {
			userTemplate.Pfp.Src = Circle(useri.Pfp)
			DrawImage(dc, &userTemplate.Pfp)
		}

		// draw XP Bar
		if (userTemplate.XpBar != XpBar{}) {
			_xpbar := userTemplate.XpBar
			_xpbar.XP = useri.XP
			_xpbar.MaxXP = useri.MaxXP
			DrawXpBar(dc, _xpbar)
		}

		// draw XPMaxXP
		if userTemplate.XPAndMaxXP != emptyTextBox {
			_XPAndMaxXP := userTemplate.XPAndMaxXP
			_XPAndMaxXP.Text = TightyNumbers(useri.XP)

			_XPAndMaxXP.TextAfter.Text += TightyNumbers(useri.MaxXP)

			DrawText(dc, &_XPAndMaxXP, defaultFontFace, defaultFontSize)
		}

		// draw XP
		if userTemplate.XP != emptyTextBox {
			_XP := userTemplate.XP
			_XP.Text = TightyNumbers(useri.XP)

			DrawText(dc, &_XP, defaultFontFace, defaultFontSize)
		}

		// draw level
		if userTemplate.Level != emptyTextBox {
			_level := userTemplate.Level
			_level.Text = TightyNumbers(useri.Level)

			DrawText(dc, &_level, defaultFontFace, defaultFontSize)
		}

		// draw MemberSince
		if userTemplate.MemberSince != emptyTextBox {
			_MemberSince := userTemplate.MemberSince
			_MemberSince.Text = useri.MemberSince

			DrawText(dc, &_MemberSince, defaultFontFace, defaultFontSize)
		}

		// draw Text XP
		if userTemplate.TextXP != emptyTextBox {
			_textXP := userTemplate.TextXP
			_textXP.Text = TightyNumbers(useri.TextXP)

			DrawText(dc, &_textXP, defaultFontFace, defaultFontSize)
		}

		// draw VOICE XP
		if userTemplate.VoiceXP != emptyTextBox {
			_voiceXP := userTemplate.VoiceXP
			_voiceXP.Text = TightyNumbers(useri.VoiceXP)

			DrawText(dc, &_voiceXP, defaultFontFace, defaultFontSize)
		}

		// draw MULTIPLIER
		if userTemplate.Multiplier != emptyTextBox {
			_multiplier := userTemplate.Multiplier
			_multiplier.Text = TightyNumbers(useri.Multiplier)

			DrawText(dc, &_multiplier, defaultFontFace, defaultFontSize)
		}

		// draw MESSAGES
		if userTemplate.Messages != emptyTextBox {
			_messages := userTemplate.Messages
			_messages.Text = TightyNumbers(useri.Messages)

			DrawText(dc, &_messages, defaultFontFace, defaultFontSize)
		}

		// draw VOICEMINUTES XP
		if userTemplate.VoiceMinutes != emptyTextBox {
			_voiceMinutes := userTemplate.VoiceMinutes
			_voiceMinutes.Text = TightyNumbers(useri.VoiceMinutes)

			DrawText(dc, &_voiceMinutes, defaultFontFace, defaultFontSize)
		}
	}

	return dc.Image()
}

func DrawXpBar(dc *gg.Context, xpBar XpBar) {
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

func DrawLevelBar(dc *gg.Context, val int, x, y, max, maxX, maxY float64) {
	dc.SetHexColor("#ffffff")
}

func DrawText(dc *gg.Context, text *TextObject, defaultFont *truetype.Font, defaultFontSize float64) {
	DefaultingFontOfTextObject(text, defaultFont, defaultFontSize)
	fontface := EasyGetFontFace(text.FontFace, text.FontSize)
	dc.SetFontFace(fontface)
	w, _ := dc.MeasureString(text.Text)

	if text.Centered {
		// draw text before
		if text.TextBefore != nil {
			text.TextBefore.X += text.X - (w / 2)
			text.TextBefore.Y += text.Y
			DrawText(dc, text.TextBefore, defaultFont, defaultFontSize)
		}

		dc.SetFontFace(fontface)
		dc.SetHexColor(text.Color)
		dc.DrawStringAnchored(text.Text, text.X, text.Y, 0.5, 0.5)

		// draw text after
		if text.TextAfter != nil {
			text.TextAfter.X += text.X + (w / 2)
			text.TextAfter.Y += text.Y
			DrawText(dc, text.TextAfter, defaultFont, defaultFontSize)
		}
	} else if text.RightAligned {
		// draw text before
		if text.TextBefore != nil {
			text.TextBefore.X += text.X - w
			text.TextBefore.Y += text.Y
			DrawText(dc, text.TextBefore, defaultFont, defaultFontSize)
		}

		dc.SetFontFace(fontface)
		dc.SetHexColor(text.Color)
		dc.DrawStringAnchored(text.Text, text.X, text.Y, 1, 0.5)

		// draw text after
		if text.TextAfter != nil {
			text.TextAfter.X += text.X
			text.TextAfter.Y += text.Y
			DrawText(dc, text.TextAfter, defaultFont, defaultFontSize)
		}
	} else {
		// draw text before
		if text.TextBefore != nil {
			text.TextBefore.X += text.X
			text.TextBefore.Y += text.Y
			DrawText(dc, text.TextBefore, defaultFont, defaultFontSize)
		}

		dc.SetFontFace(fontface)
		dc.SetHexColor(text.Color)
		dc.DrawStringAnchored(text.Text, text.X, text.Y, 0, 0.5)

		// draw text after
		if text.TextAfter != nil {
			text.TextAfter.X += text.X + w
			text.TextAfter.Y += text.Y
			DrawText(dc, text.TextAfter, defaultFont, defaultFontSize)
		}
	}
}

func DrawImage(dc *gg.Context, image *ImageObject) {
	if image.Centered {
		dc.DrawImageAnchored(resize.Resize(image.W, image.H, Circle(image.Src), resize.Lanczos2), image.X, image.Y, 0.5, 0.5)
	} else if image.RightAligned {
		dc.DrawImageAnchored(resize.Resize(image.W, image.H, Circle(image.Src), resize.Lanczos2), image.X, image.Y, 1, 0.5)
	} else {
		dc.DrawImageAnchored(resize.Resize(image.W, image.H, Circle(image.Src), resize.Lanczos2), image.X, image.Y, 0, 0)
	}
}
