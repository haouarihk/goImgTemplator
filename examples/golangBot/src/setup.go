package main

import (
	"image"
	"os"

	Templator "github.com/haouarihk/goImgTemplator/src"
)

func setupThemes() map[string]Templator.Theme {
	defaultFont := Templator.EasyLoadFontFace("fonts/SourceSansPro-Bold.ttf")

	// levelup page
	levelUpTemplate, _ := os.Open("templates/levelUp.png")
	defer levelUpTemplate.Close()

	levelUpTemplateShiny, _ := os.Open("templates/levelUpShiny.png")
	defer levelUpTemplateShiny.Close()

	levelUpTemplateSrc, _, _ := image.Decode(levelUpTemplate)
	levelUpTemplateShinySrc, _, _ := image.Decode(levelUpTemplateShiny)

	// levely page
	levelyTemplate, _ := os.Open("templates/levely.png")
	defer levelyTemplate.Close()

	levelyTemplateShiny, _ := os.Open("templates/levelyShiny.png")
	defer levelyTemplateShiny.Close()

	levelyTemplateSrc, _, _ := image.Decode(levelyTemplate)
	levelyTemplateShinySrc, _, _ := image.Decode(levelyTemplateShiny)

	// leaderboard page
	leaderboardTemplate, _ := os.Open("templates/leaderboard.png")
	defer leaderboardTemplate.Close()

	leaderboardTemplateShiny, _ := os.Open("templates/leaderboardShiny.png")
	defer leaderboardTemplateShiny.Close()

	leaderboardTemplateSrc, _, _ := image.Decode(leaderboardTemplate)
	leaderboardTemplateShinySrc, _, _ := image.Decode(leaderboardTemplateShiny)

	return map[string]Templator.Theme{
		"default": {
			DefaultFontFace: defaultFont,
			DefaultFontSize: 20,
			Packs: map[string]Templator.ThemePack{
				"levelUp": {
					TemplateSrc:      levelUpTemplateSrc,
					TemplateSrcShiny: levelUpTemplateShinySrc,
					Scale:            2,

					// this is the template for the user properties, such as his pfp, username, etc
					UserTemplate: []Templator.UserTemplate{
						{
							Pfp: Templator.ImageObject{
								// this should be the default pfp
								Src: nil,
								// location of the pfp
								X: 129,
								Y: 162,
								// size of the pfp
								W: 351,
								H: 351,
							},
							FullUsername: Templator.TextObject{
								// this should be the default username
								Text: "user",
								// text color
								Color: "#04A8C3",
								// location of the text
								X: 900,
								Y: 162,
								// font settings
								FontFace: defaultFont,
								FontSize: 72,
								Centered: true,
							},
							Level: Templator.TextObject{
								Text:     "LEVEL",
								Color:    "#04A8C3",
								X:        1150,
								Y:        400,
								FontFace: defaultFont,
								FontSize: 48,
								Centered: true,
							},
						},
					},

					// all of these texts are staticaly
					Texts: []Templator.TextObject{
						{
							// this should be the default username
							Text: "CONGRATULATIONS !",
							// text color
							Color: "#bbbbbb",
							// location of the text
							X: 900,
							Y: 70,
							// font settings
							FontFace: defaultFont,
							FontSize: 52,
							Centered: true,
						},
						{
							Text:     "LEVEL",
							Color:    "#bbbbbb",
							X:        1150,
							Y:        340,
							FontFace: defaultFont,
							FontSize: 52,
							Centered: true,
						},
						{
							Text:     "YOU LEVELED UP !",
							Color:    "#212129",
							X:        880,
							Y:        550,
							FontFace: defaultFont,
							FontSize: 74,
							Centered: true,
						},
					},
				},
				"levely": {
					TemplateSrc:      levelyTemplateSrc,
					TemplateSrcShiny: levelyTemplateShinySrc,
					Scale:            1,
					UserTemplate: []Templator.UserTemplate{
						{
							Pfp: Templator.ImageObject{
								// this should be the default pfp
								Src: nil,
								// location of the pfp
								X: 34,
								Y: 32,
								// size of the pfp
								W: 516,
								H: 516,
							},
							FullUsername: Templator.TextObject{
								// this should be the default username
								Text: "USERNAME#0001",
								// text color
								Color: "#212129",
								// location of the text
								X: 970,
								Y: 200,
								// font settings
								FontSize: 79,
								Centered: true,
							},
							MemberSince: Templator.TextObject{
								Text:     "111",
								Color:    "#04A8C3",
								X:        1300,
								Y:        374,
								FontSize: 66,
								Centered: true,
							},
							Level: Templator.TextObject{
								Text:     "111",
								Color:    "#04A8C3",
								X:        1300,
								Y:        620,
								FontSize: 92,
								Centered: true,
								TextBefore: &Templator.TextObject{
									Text:  "Level",
									Color: "#BBBBBB",

									// offset x,y from the parent
									X:            -10,
									Y:            10,
									FontSize:     64,
									RightAligned: true,
								},
							},
							XPAndMaxXP: Templator.TextObject{
								Text:     "111",
								Color:    "#04A8C3",
								X:        100,
								Y:        620,
								FontSize: 92,
								TextAfter: &Templator.TextObject{
									// i thought differently on this one
									// whatever you type here, will be between the two texts
									Text: "/",

									Color:    "#BBBBBB",
									X:        10,
									Y:        10,
									FontSize: 64,
									TextAfter: &Templator.TextObject{
										Text:     "XP",
										Color:    "#04A8C3",
										X:        10,
										Y:        -10,
										FontSize: 92,
									},
								},
							},
							XpBar: Templator.XpBar{
								X:          90,
								Y:          688.8,
								Width:      1397,
								Height:     70,
								Roundyness: 30,
								Color:      "#04A8C3",
							},
							TextXP: Templator.TextObject{
								Text:     "2",
								Color:    "#04A8C3",
								X:        1780,
								Y:        160,
								FontSize: 72,
								Centered: true,
							},
							VoiceXP: Templator.TextObject{
								Text:     "2",
								Color:    "#04A8C3",
								X:        1780,
								Y:        396,
								FontSize: 72,
								Centered: true,
							},
							Multiplier: Templator.TextObject{
								Text:         "2",
								Color:        "#04A8C3",
								X:            2110,
								Y:            720,
								FontSize:     76,
								RightAligned: true,
								TextAfter: &Templator.TextObject{
									Text:     "X",
									Color:    "#04A8C3",
									X:        10,
									Y:        0,
									FontSize: 64,
								},
							},
							Messages: Templator.TextObject{
								Text:     "2",
								Color:    "#04A8C3",
								X:        2096,
								Y:        920,
								FontSize: 62,
								Centered: true,
							},
							VoiceMinutes: Templator.TextObject{
								Text:     "2",
								Color:    "#04A8C3",
								X:        2096,
								Y:        1120,
								FontSize: 62,
								Centered: true,
							},
						},
					},

					// all of these texts are staticaly
					Texts: []Templator.TextObject{
						{
							Text: "TEXT XP",
							// text color
							Color: "#bbbbbb",
							// location of the text
							X: 2020,
							Y: 160,
							// font settings
							FontSize: 72,
						},
						{
							Text: "VOICE XP",
							// text color
							Color: "#bbbbbb",
							// location of the text
							X: 2020,
							Y: 396,
							// font settings
							FontSize: 72,
						},
						{
							Text:     "MULTIPLIER",
							Color:    "#212129",
							X:        2240,
							Y:        720,
							FontSize: 72,
						},
						{
							Text:     "MESSAGES",
							Color:    "#212129",
							X:        2240,
							Y:        920,
							FontSize: 72,
						},
						{
							Text:     "VOICE",
							Color:    "#212129",
							X:        2240,
							Y:        1080,
							FontSize: 72,
						},
						{
							Text:     "MINUTES",
							Color:    "#212129",
							X:        2240,
							Y:        1150,
							FontSize: 72,
						},
						{
							Text:     "MEMBER SINCE",
							Color:    "#bbbbbb",
							X:        576,
							Y:        378,
							FontSize: 72,
						},
					},
				},
				"leaderboard": {
					TemplateSrc:      leaderboardTemplateSrc,
					TemplateSrcShiny: leaderboardTemplateShinySrc,
					// this is the template for the user properties, such as his pfp, username, etc
					UserTemplate: []Templator.UserTemplate{
						{
							Pfp: Templator.ImageObject{
								// this should be the default pfp
								Src: nil,
								// location of the pfp
								X: 675,
								Y: 164,
								// size of the pfp
								W:        255,
								H:        255,
								Centered: true,
							},
							FullUsername: Templator.TextObject{
								// this should be the default username
								Text: "USERNAME#0001",
								// text color
								Color: "#04A8C3",
								// location of the text
								X: 675,
								Y: 310,
								// font settings
								FontSize: 46,
								Centered: true,
							},
							Level: Templator.TextObject{
								Text:         "111",
								Color:        "#04A8C3",
								X:            675,
								Y:            360,
								FontSize:     36,
								RightAligned: true,
								TextBefore: &Templator.TextObject{
									Text:         "Level",
									Color:        "#bbbbbb",
									X:            -10,
									Y:            0,
									FontSize:     36,
									RightAligned: true,
								},
							},
							XP: Templator.TextObject{
								Text:     "2",
								Color:    "#04A8C3",
								X:        750,
								Y:        360,
								FontSize: 36,
								TextBefore: &Templator.TextObject{
									Text:         "XP",
									Color:        "#bbbbbb",
									X:            -10,
									Y:            0,
									FontSize:     36,
									RightAligned: true,
								},
							},
						},
						{
							Pfp: Templator.ImageObject{
								// this should be the default pfp
								Src: nil,
								// location of the pfp
								X: 272,
								Y: 283,
								// size of the pfp
								W:        255,
								H:        255,
								Centered: true,
							},
							FullUsername: Templator.TextObject{
								// this should be the default username
								Text: "USERNAME#0001",
								// text color
								Color: "#04A8C3",
								// location of the text
								X: 272,
								Y: 430,
								// font settings
								FontSize: 46,
								Centered: true,
							},
							Level: Templator.TextObject{
								Text:         "111",
								Color:        "#04A8C3",
								X:            272,
								Y:            470,
								FontSize:     36,
								RightAligned: true,
								TextBefore: &Templator.TextObject{
									Text:         "Level",
									Color:        "#bbbbbb",
									X:            -10,
									Y:            0,
									FontSize:     36,
									RightAligned: true,
								},
							},
							XP: Templator.TextObject{
								Text:     "2",
								Color:    "#04A8C3",
								X:        350,
								Y:        470,
								FontSize: 36,
								TextBefore: &Templator.TextObject{
									Text:         "XP",
									Color:        "#bbbbbb",
									X:            -10,
									Y:            0,
									FontSize:     36,
									RightAligned: true,
								},
							},
						},
						{
							Pfp: Templator.ImageObject{
								// this should be the default pfp
								Src: nil,
								// location of the pfp
								X: 1074,
								Y: 353,
								// size of the pfp
								W:        250,
								H:        250,
								Centered: true,
							},
							FullUsername: Templator.TextObject{
								// this should be the default username
								Text: "USERNAME#0001",
								// text color
								Color: "#04A8C3",
								// location of the text
								X: 1074,
								Y: 495,
								// font settings
								FontSize: 46,
								Centered: true,
							},
							Level: Templator.TextObject{
								Text:         "111",
								Color:        "#04A8C3",
								X:            1074,
								Y:            550,
								FontSize:     36,
								RightAligned: true,
								TextBefore: &Templator.TextObject{
									Text:         "Level",
									Color:        "#bbbbbb",
									X:            -10,
									Y:            0,
									FontSize:     36,
									RightAligned: true,
								},
							},
							XP: Templator.TextObject{
								Text:     "2",
								Color:    "#04A8C3",
								X:        1140,
								Y:        550,
								FontFace: defaultFont,
								FontSize: 36,
								TextBefore: &Templator.TextObject{
									Text:         "XP",
									Color:        "#bbbbbb",
									X:            -10,
									Y:            0,
									FontSize:     36,
									RightAligned: true,
								},
							},
						},
					},

					// key words for the text: username, level, xp
					// any other text will not be dynamically changed
					Texts: []Templator.TextObject{
						{
							Text: "MONTH",
							// text color
							Color: "#bbbbbb",
							// location of the text
							X: 1010,
							Y: 55,
							// font settings
							FontSize: 52,
							TextAfter: &Templator.TextObject{
								Text: "LY",
								// text color
								Color: "#04A8C3",
								// position relative to the parent
								X: 0,
								Y: 0,
								// font settings
								FontSize: 52,
							},
						},
					},
				},
			},
		},
	}
}
