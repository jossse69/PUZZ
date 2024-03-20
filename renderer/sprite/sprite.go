package sprite

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jossse69/PUZZ/commons"
)

type Sprite struct {
	Pixels []int
	Width  int
	Height int
}

// NewSprite creates a new sprite with the given width and height.
func NewSprite(width, height int) *Sprite {
	return &Sprite{
		Pixels: make([]int, width*height),
		Width:  width,
		Height: height,
	}
}

// SetPixel sets the color ID of a pixel in the sprite.
func (s *Sprite) SetPixel(x, y, colorID int) {
	if x < 0 || y < 0 || x >= s.Width || y >= s.Height {
		return // Out of bounds
	}
	s.Pixels[y*s.Width+x] = colorID
}

// GetPixel returns the color ID of a pixel in the sprite.
func (s *Sprite) GetPixel(x, y int) int {
	if x < 0 || y < 0 || x >= s.Width || y >= s.Height {
		return -1 // Out of bounds or considered transparent
	}
	return s.Pixels[y*s.Width+x]
}

// Crop crops the sprite with the given x, y, width, and height. returns a new sprite.
func (s *Sprite) Crop(x, y, width, height int) *Sprite {
	cropped := NewSprite(width, height)
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			colorID := s.GetPixel(x+i, y+j)
			cropped.SetPixel(i, j, colorID)
		}
	}
	return cropped
}

// LoadFromImage loads a sprite from an image by the filepath.
func LoadFromImage(path string) *Sprite {

	// get the image located at path
	img := rl.LoadImage(path)
	// create a new sprite with the width and height of the image
	s := NewSprite(int(img.Width), int(img.Height))
	// load the image into the sprite
	for x := 0; x < int(img.Width); x++ {
		for y := 0; y < int(img.Height); y++ {
			// convert the colors of the image to the closest colors in the pallet
			s.SetPixel(x, y, int(GetClosestColor(rl.GetImageColor(*img, int32(x), int32(y)))))
		}
	}
	return s

}

// GetClosestColor returns the closest color in the pallet to the given color.
func GetClosestColor(color rl.Color) int {
	var closestColor int
	var minDistance float64 = math.MaxFloat64

	// if the color is transparent, return -1
	if color.A == 0 {
		return -1
	}

	for i, paletteColor := range commons.Colors {
		// Extract the RGB components of the palette color
		paletteR := uint8((paletteColor & 0xFF0000) >> 16)
		paletteG := uint8((paletteColor & 0x00FF00) >> 8)
		paletteB := uint8(paletteColor & 0x0000FF)

		// Calculate the Euclidean distance between the given color and the palette color
		distance := math.Sqrt(math.Pow(float64(color.R)-float64(paletteR), 2) +
			math.Pow(float64(color.G)-float64(paletteG), 2) +
			math.Pow(float64(color.B)-float64(paletteB), 2))

		// Update the closest color if the distance is smaller
		if distance < minDistance {
			minDistance = distance
			closestColor = i
		}
	}

	return closestColor
}
