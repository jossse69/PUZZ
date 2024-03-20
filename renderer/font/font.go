package font

import (
	"github.com/jossse69/PUZZ/renderer/sprite"
)

type Font struct {
	Sprite      *sprite.Sprite // The sprite containing the font's characters
	CharWidth   int            // Width of a single character
	CharHeight  int            // Height of a single character
	CharsPerRow int            // Number of characters per row in the sprite
	FirstChar   rune           // The first character represented in the sprite (often ' ')
}

// NewFont creates a new Font instance from a sprite.
func NewFont(s *sprite.Sprite, charWidth, charHeight, charsPerRow int, firstChar rune) *Font {
	return &Font{
		Sprite:      s,
		CharWidth:   charWidth,
		CharHeight:  charHeight,
		CharsPerRow: charsPerRow,
		FirstChar:   firstChar,
	}
}

// LoadFontFromImage loads a font from an image file.
func LoadFontFromImage(path string, charWidth, charHeight, charsPerRow int, firstChar rune) *Font {
	fontSprite := sprite.LoadFromImage(path)
	return NewFont(fontSprite, charWidth, charHeight, charsPerRow, firstChar)
}
