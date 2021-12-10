package Templator

import (
	"image"

	"github.com/golang/freetype/truetype"
)

/* This is the image Object, tells us about the location, the image itself and width/height. For costumization */
type ImageObject struct {
	/* The source image */
	Src image.Image

	/* Possition */
	X, Y int

	/* Size */
	W, H uint

	/* Is it Centered */
	Centered bool

	/* Is it Right Aligned, can't be used with Centered */
	RightAligned bool
}

/* This is the text Object, tells us about the location, the text itself and width/height. For costumization */
type TextObject struct {
	/* text content */
	Text string

	/* color of the text */
	Color string

	/* Possition relative to the parent*/
	X, Y float64

	/* Font Face */
	FontFace *truetype.Font

	/* Font Size */
	FontSize float64

	/* The text that comes after this(it will became the parent) */
	TextAfter *TextObject

	/* The text that comes before this(it will became the parent) */
	TextBefore *TextObject

	/* Is it Centered */
	Centered bool
	/* Is it Right Aligned, can't be used with Centered */
	RightAligned bool
}

/* The way the user properties displayed in the template */
type UserTemplate struct {
	/* Template of: The Profile image of the user */
	Pfp ImageObject

	/* Template of: The username of the user, (without the tag)  */
	Username TextObject

	/* Template of: The tag of the user */
	Tag TextObject

	/* Template of: The username and tag of the user */
	FullUsername TextObject

	/* Template of: The username and tag being the TextAfter */
	TagTextAfterUsername TextObject

	/* Template of: The Level of the user */
	Level TextObject

	/* Template of: The XP of the user */
	XP TextObject

	/* this is a special case where you want the MaxXP to display after The XP, the MaxXp gonna be the textAfter, any text in the textAfter will be displayed at the begining of textAfter. */
	XPAndMaxXP TextObject

	/* Template of: The XP bar */
	XpBar XpBar

	/* Template of: MaxXP of the user*/
	MaxXP TextObject

	/* Template of: The Date on when the user became a member (string)*/
	MemberSince TextObject

	/* Template of: The Text XP of the user */
	TextXP TextObject

	/* Template of: The Voice XP of the user */
	VoiceXP TextObject

	/* Template of: The Multiplier of the user's XP */
	Multiplier TextObject

	/* Template of: The Messages that the user has sent */
	Messages TextObject

	/* Template of: The amount of minutes the user has been in the vc*/
	VoiceMinutes TextObject
}

/* This is the user input, tells us all about the user, his username, level, xp ... etc */
type UserInput struct {
	/* The Profile image of the user */
	Pfp image.Image

	/* The username of the user */
	FullUsername string

	/* The Level of the user */
	Level int

	/* The XP of the user */
	XP int

	/* Template of: how much xp is needed to level up from this level */
	MaxXP int

	/* The Date on when the user became a member (string)*/
	MemberSince string

	/* The Text XP of the user */
	TextXP int

	/* The Voice XP of the user */
	VoiceXP int

	/* The Multiplier of the user's XP */
	Multiplier int

	/* The Messages that the user has sent */
	Messages int

	/* The amount of minutes the user has been in the vc*/
	VoiceMinutes int
}

/* This is the additional options for the render function, tells us about other stuff that are not related to the user directly, such as, shiny, theme, ConstumeBackground... etc */
type Options struct {
	/* Use the shiny version of the template */
	Shiny bool

	/* A costume background set by the user*/
	CostumeBackground image.Image

	/* A costume font set by the user*/
	ConstumeFontFace *truetype.Font

	/* The used theme */
	Theme string

	/* the used Pack (levely, levelUp, leaderboard ...etc), you can add costume ones to the theme if you want */
	Pack string
}

/* This is the Theme pack, it's the pack. (ex: levely, levelUp... etc) */
type ThemePack struct {
	/* The Image of the template */
	TemplateSrc image.Image

	/* The Image of the template but shiny (optional)*/
	TemplateSrcShiny image.Image

	/* The Scale of the image, recommended: X1, X2 (MAX)*/
	Scale int

	/* A static text displayed in the template (follows the font choice)*/
	Texts []TextObject

	/* A static images displayed in the template (optional) */
	images []ImageObject

	/* The way the user properties displayed in the template */
	UserTemplate []UserTemplate
}

/* This is the theme, that holds to the packs and some default stuff */
type Theme struct {
	// default font path
	DefaultFontFace *truetype.Font
	DefaultFontSize float64

	// Packs
	Packs map[string]ThemePack
}

/* The class for the main object */
type Templater struct {
	Themes map[string]Theme
}

/* This is the XpBar */
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
