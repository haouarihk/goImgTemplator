package Templator

import (
	"fmt"
	"image"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/nfnt/resize"
)

func Init(themes map[string]Theme) Templater {
	return Templater{
		Themes: themes,
	}
}

var emptyTextBox = TextObject{}

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
		drawText(dc, &text, defaultFontFace, defaultFontSize)
	}

	// ---- draw user properties
	for i, userTemplate := range thatPack.UserTemplate {
		useri := user[i]

		// draw username
		drawUsername(dc, userTemplate, useri, defaultFontFace, defaultFontSize)

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

		// draw xp and Max XP
		drawXPMaxXP(dc, userTemplate, useri, defaultFontFace, defaultFontSize)

		// draw level
		if userTemplate.Level != emptyTextBox {
			_level := userTemplate.Level
			_level.Text = TightyNumbers(useri.Level)

			drawText(dc, &_level, defaultFontFace, defaultFontSize)
		}

		// draw MemberSince
		if userTemplate.MemberSince != emptyTextBox {
			_MemberSince := userTemplate.MemberSince
			_MemberSince.Text = useri.MemberSince

			drawText(dc, &_MemberSince, defaultFontFace, defaultFontSize)
		}

		// draw Text XP
		if userTemplate.TextXP != emptyTextBox {
			_textXP := userTemplate.TextXP
			_textXP.Text = TightyNumbers(useri.TextXP)

			drawText(dc, &_textXP, defaultFontFace, defaultFontSize)
		}

		// draw VOICE XP
		if userTemplate.VoiceXP != emptyTextBox {
			_voiceXP := userTemplate.VoiceXP
			_voiceXP.Text = TightyNumbers(useri.VoiceXP)

			drawText(dc, &_voiceXP, defaultFontFace, defaultFontSize)
		}

		// draw MULTIPLIER
		if userTemplate.Multiplier != emptyTextBox {
			_multiplier := userTemplate.Multiplier
			_multiplier.Text = TightyNumbers(useri.Multiplier)

			drawText(dc, &_multiplier, defaultFontFace, defaultFontSize)
		}

		// draw MESSAGES
		if userTemplate.Messages != emptyTextBox {
			_messages := userTemplate.Messages
			_messages.Text = TightyNumbers(useri.Messages)

			drawText(dc, &_messages, defaultFontFace, defaultFontSize)
		}

		// draw VOICEMINUTES XP
		if userTemplate.VoiceMinutes != emptyTextBox {
			_voiceMinutes := userTemplate.VoiceMinutes
			_voiceMinutes.Text = TightyNumbers(useri.VoiceMinutes)

			drawText(dc, &_voiceMinutes, defaultFontFace, defaultFontSize)
		}
	}

	return dc.Image()
}

func drawUsername(dc *gg.Context, userTemplate UserTemplate, useri UserInput, defaultFontFace *truetype.Font, defaultFontSize float64) {
	username, tag, err := SeperateUsernameFromTag(useri.FullUsername)

	if err != "" {
		fmt.Println(err, useri.FullUsername)
		return
	}

	// draw fullUsername
	if userTemplate.FullUsername != emptyTextBox {
		_username := userTemplate.FullUsername
		_username.Text = useri.FullUsername

		drawText(dc, &_username, defaultFontFace, defaultFontSize)
	}

	// draw usernameSeperateFromTag
	if userTemplate.TagTextAfterUsername != emptyTextBox {
		_username := userTemplate.TagTextAfterUsername
		_username.Text = useri.FullUsername

		if userTemplate.TagTextAfterUsername.TextAfter != nil {
			_username.TextAfter.Text += tag
		}

		drawText(dc, &_username, defaultFontFace, defaultFontSize)
	}

	// draw user's tag
	if userTemplate.Tag != emptyTextBox {
		_username := userTemplate.Tag
		_username.Text = tag

		drawText(dc, &_username, defaultFontFace, defaultFontSize)
	}

	// draw just the username
	if userTemplate.Username != emptyTextBox {
		_username := userTemplate.Username
		_username.Text = username

		drawText(dc, &_username, defaultFontFace, defaultFontSize)
	}
}

func drawXPMaxXP(dc *gg.Context, userTemplate UserTemplate, useri UserInput, defaultFontFace *truetype.Font, defaultFontSize float64) {
	// draw XPMaxXP
	if userTemplate.XPAndMaxXP != emptyTextBox {
		_XPAndMaxXP := userTemplate.XPAndMaxXP
		_XPAndMaxXP.Text = TightyNumbers(useri.XP)

		_XPAndMaxXP.TextAfter.Text += TightyNumbers(useri.MaxXP)

		drawText(dc, &_XPAndMaxXP, defaultFontFace, defaultFontSize)
	}

	// draw XP
	if userTemplate.XP != emptyTextBox {
		_XP := userTemplate.XP
		_XP.Text = TightyNumbers(useri.XP)

		drawText(dc, &_XP, defaultFontFace, defaultFontSize)
	}

	// draw Max XP
	if userTemplate.MaxXP != emptyTextBox {
		_MaxXP := userTemplate.MaxXP
		_MaxXP.Text = TightyNumbers(useri.MaxXP)

		drawText(dc, &_MaxXP, defaultFontFace, defaultFontSize)
	}
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
		dc.SetFillStyle(gg.NewSolidPattern(ConvertToHex(xpBar.Color)))
		dc.DrawRoundedRectangle(xpBar.X, xpBar.Y, leni+14, xpBar.Height, xpBar.Roundyness)
		dc.Fill()
	}
}

func drawText(dc *gg.Context, text *TextObject, defaultFont *truetype.Font, defaultFontSize float64) {
	DefaultingFontOfTextObject(text, defaultFont, defaultFontSize)
	fontface := EasyGetFontFace(text.FontFace, text.FontSize)
	dc.SetFontFace(fontface)
	w, _ := dc.MeasureString(text.Text)

	if text.Centered {
		// draw text before
		if text.TextBefore != nil {
			text.TextBefore.X += text.X - (w / 2)
			text.TextBefore.Y += text.Y
			drawText(dc, text.TextBefore, defaultFont, defaultFontSize)
		}

		dc.SetFontFace(fontface)
		dc.SetHexColor(text.Color)
		dc.DrawStringAnchored(text.Text, text.X, text.Y, 0.5, 0.5)

		// draw text after
		if text.TextAfter != nil {
			text.TextAfter.X += text.X + (w / 2)
			text.TextAfter.Y += text.Y
			drawText(dc, text.TextAfter, defaultFont, defaultFontSize)
		}
	} else if text.RightAligned {
		// draw text before
		if text.TextBefore != nil {
			text.TextBefore.X += text.X - w
			text.TextBefore.Y += text.Y
			drawText(dc, text.TextBefore, defaultFont, defaultFontSize)
		}

		dc.SetFontFace(fontface)
		dc.SetHexColor(text.Color)
		dc.DrawStringAnchored(text.Text, text.X, text.Y, 1, 0.5)

		// draw text after
		if text.TextAfter != nil {
			text.TextAfter.X += text.X
			text.TextAfter.Y += text.Y
			drawText(dc, text.TextAfter, defaultFont, defaultFontSize)
		}
	} else {
		// draw text before
		if text.TextBefore != nil {
			text.TextBefore.X += text.X
			text.TextBefore.Y += text.Y
			drawText(dc, text.TextBefore, defaultFont, defaultFontSize)
		}

		dc.SetFontFace(fontface)
		dc.SetHexColor(text.Color)
		dc.DrawStringAnchored(text.Text, text.X, text.Y, 0, 0.5)

		// draw text after
		if text.TextAfter != nil {
			text.TextAfter.X += text.X + w
			text.TextAfter.Y += text.Y
			drawText(dc, text.TextAfter, defaultFont, defaultFontSize)
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
