# this is the Documentation for the project
This engine is a class, you need to initialize it with: themes.


## Notes:
pack: is the page (ex: levelUp, leaderboard, levely).
theme: is the collection of packs.



## How to start using it:

```go
templator := Templator.Templater{
		Themes: setupThemes(),
}
```

and then you can use the function render:

```go
templator.Render(
// the users passed in
[]Templator.UserTemplate{{
    Username:    "username",
    		Tag:         "#0001",
    		Pfp:         UIMGSrc,
    		Level:       1000000,
    		XP:          50000000,
    		MemberSince: "01/10/01",
    }
}, {
// the additional options

// is it shiny
Shiny:             false,

// costume background: image.Image
CostumeBackground: kSrc,

// theme name
Theme:             "default",

// the pack (ex: levely, leaderboard, etc)
Pack:              "leaderboard",

})
```

The reason why the first argument is an array and not one, is because you can have multiple users in the leaderboard pack.



## How to make themes:
First of all you need to have a template to base everything around it
the template to be consistant of multiple packs.


Then you take a big breath and start with making the setupThemes function.
```go
func setupThemes(){
    // load all the images from the templates folder
    levelUpTemplate, _ := os.Open("templates/levelUp.png")
	defer levelUpTemplate.Close()
    levelUpTemplateShiny, _ := os.Open("templates/levelUpShiny.png")
	defer levelUpTemplateShiny.Close()

    // decode them into images.Image
    levelUpTemplateSrc, _, _ := image.Decode(levelUpTemplate)
	levelUpTemplateShinySrc, _, _ := image.Decode(levelUpTemplateShiny)

    // setup a defaultFont
    defaultFont := Templator.EasyLoadFontFace("fonts/SourceSansPro-Bold.ttf")

    // then you can start making the themes
    return 	return map[string]Templator.Theme{
		"default": {
            // this is the font that every text will be based on, unless, the user specifies a font or you override it.
			DefaultFontFace:   defaultFont,
			DefaultFontSize:   20,

            // these are the list of packs in this template
			Packs: map[string]Templator.ThemePack{
                {
                "levelUp": {
                    // the template image
					TemplateSrc:      levelUpTemplateSrc,

                    // the template image when it's shiny
					TemplateSrcShiny: levelUpTemplateShinySrc,

                    // the scale of the overall image (it only effect performance)
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

							Username: Templator.TextObject{
								// this should be the default username
								Text: "user",

								// text color
								Color: "#04A8C3",

								// location of the text
								X: 900,
								Y: 162,

								// (optional): it will override the user costume font settings
								FontFace: defaultFont,
								FontSize: 72,

                                // centering the text
								Centered: true,

								TextAfter: &Templator.TextObject{
									// this should be the default tag
									Text: "#0001",
									// text color
									Color: "#04A8C3",
									// font settings

                                    // (optional): it will override the user costume font settings
									FontFace: defaultFont,
									FontSize: 72,
								},
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
					Texts: []TextObject{
						{
							// this should be the default username
							Text: "CONGRATULATIONS !",

							// text color
							Color: "#bbbbbb",

							// location of the text
							X: 900,
							Y: 70,

							// (optional): it will override the user costume font settings
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
                }
            }
        }
    }


}
```


### TextObject:
this is a very handy peace of code that you can use to really manipulate the text, the way you want (honestly spent too much time designing it).


You know that manipulating text in a pixelated envirement is pretty weird,
imagine the text being too big or too small, and you want to put a text infront of it, and dynamically position it bassed on the content of that text, pretty weird mir?

That's why i made this handy tool

it has
```go
// this is the text that's gonna be there (if it's in the UserTemplate, it will be the default text)
Text: "USERNAME#0001",

// this is the color of the text
Color: "#212129",

// the position of the text
X: 970,
Y: 200,

// font settings
FontFace: defaultFont,
FontSize: 72,
// if you want to center the text
Centered: true,
// if you want to rightAlign the text (note you can use eigther and not both)
// RightAlign: true,

// now to the fun part
// this is the text that's gonna be infront of it (optional)
TextAfter: &TextObject{
    // the text
    Text: "#0001",
    // text color
    Color: "#212129",
    // font settings
    FontFace: defaultFont,
    FontSize: 72,

    // the position of this text, Note that it's relative to the parent 
    X:0,
    Y:0,

    // you can also center or rightAlign, but it's in the bleeding edge.


    // you can also take things to the extreme and add another TextAfter
    // TextAfter: ...
}

// this is the text that's gonna be behind it (optional)
TextBefore: &TextObject{
    // the text
    Text: "Ye:",
    // text color
    Color: "#212129",
    // font settings
    FontFace: defaultFont,
    FontSize: 72,

    // the position of this text, Note that it's relative to the parent 
    X:0,
    Y:0,

    // you can also center or rightAlign, but it's in the bleeding edge.

    // you can also take things to the extreme and add another TextAfter
    // TextBefore: ...
}
```


### ImageObject:
As good as it sounds, it's not that advanced,
it just as the position and size, that's all




### More Notes:

- if you find your self can't center the username and tag together, you can just put the tag in the username and ditch the tag property.

- work with whatever you have, that's the maximum what i can help with, if you have any suggestions, feel free to contact me.

- I Thought differntly about XpAndMaxXp, if you put something on the Text, it will be between the parent text and the value recived from the user

