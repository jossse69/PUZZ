package renderer

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jossse69/PUZZ/commons"
	"github.com/jossse69/PUZZ/renderer/sprite"
)

type Renderer struct {
	Target *rl.RenderTexture2D
}

const (
	screenWidth  = 800
	screenHeight = 600
)

// NewRenderer creates a new renderer for the game.
//
// It takes the game width, game height, and title as parameters.
// Returns a pointer to the renderer.
func NewRenderer(gameWidth, gameHeight int, title string) *Renderer {

	rl.InitWindow(screenWidth, screenHeight, title)
	rl.SetTargetFPS(60)

	target := rl.LoadRenderTexture(int32(gameWidth), int32(gameHeight))
	rl.SetTextureFilter(target.Texture, rl.FilterPoint)

	return &Renderer{
		Target: &target,
	}
}

// StartDrawing is a function that starts the drawing process.
//
// No parameters.
// No return types.
func (r *Renderer) StartDrawing() {
	rl.BeginTextureMode(*r.Target) // Start drawing to the render texture
	rl.ClearBackground(rl.Black)   // Clear render texture background
}

// DrawOnScreen draws the renderer on the screen.
//
// No parameters.
// No return types.
func (r *Renderer) DrawOnScreen() {
	rl.EndTextureMode() // End drawing to the render texture

	rl.BeginDrawing()            // Start drawing to the main window
	rl.ClearBackground(rl.Black) // Optional: Clear main window background
	gameWidth := float32(r.Target.Texture.Width)
	gameHeight := float32(r.Target.Texture.Height)
	gameHeightScale := -screenHeight / gameHeight // for keeping the aspect ratio while scaling the texture
	gameWidthScale := screenWidth / gameWidth     // same as above
	// draw the texture on the main window
	rl.DrawTexturePro(
		r.Target.Texture,
		rl.NewRectangle(0, 0, gameWidth, gameHeight),
		rl.NewRectangle(0, -gameHeight*gameHeightScale, gameWidth*gameWidthScale, gameHeight*gameHeightScale),
		rl.NewVector2(0, 0),
		0,
		rl.White,
	)
	rl.EndDrawing() // End drawing to the main window
}

// PalletIndexToColor converts a pallet index to a color.
//
// It takes the index of the color in the pallet.
// Returns the corresponding color.
func PalletIndexToColor(index int) rl.Color {

	// create a color from the index
	selectedColor := commons.Colors[index]
	color := rl.NewColor(uint8(selectedColor>>16), uint8(selectedColor>>8), uint8(selectedColor), 255)

	return color

}

// Close releases the render texture and closes the window.
//
// No parameters.
// No return types.
func (r *Renderer) Close() {
	rl.UnloadRenderTexture(*r.Target)
	rl.CloseWindow()
}

// methods to draw stuff like text or shapes

// DrawSprite draws a sprite in the texture.
//
// It takes the x, y, and the sprite as parameters.
// No return types.
func (r *Renderer) DrawSprite(x, y int, spr *sprite.Sprite) {
	for j := 0; j < spr.Height; j++ {
		for i := 0; i < spr.Width; i++ {
			colorID := spr.GetPixel(i, j)
			if colorID != -1 { // -1 is transparent
				r.DrawPixel(x+i, y+j, colorID)
			}
		}
	}
}

// Fill fills the screen with a color.
//
// It takes the color as a parameter.
// No return types.
func (r *Renderer) Fill(color int) {
	for x := 0; x < screenWidth; x++ {
		for y := 0; y < screenHeight; y++ {
			r.DrawPixel(x, y, color)
		}
	}
}

func (r *Renderer) PrintText(x, y int, text string, fg int, bg int) {
	// draw sprites of text fonts here
}

// DrawRectangle draws a rectangle in the texture.
//
// It takes the x, y, width, height, and color as parameters.
// No return types.
func (r *Renderer) DrawRectangle(x, y, width, height int, color int) {

	// draw a rectangle in the texture
	for i := x; i < x+width; i++ {
		for j := y; j < y+height; j++ {
			r.DrawPixel(i, j, color)
		}
	}
}

// DrawRectangleLines draws a rectangle in the texture.
//
// It takes the x, y, width, height, and color as parameters.
// No return types.
func (r *Renderer) DrawRectangleLines(x, y, width, height int, color int) {
	// draw a rectangle in the texture
	for i := x; i < x+width; i++ {
		for j := y; j < y+height; j++ {
			if i == x || i == x+width-1 || j == y || j == y+height-1 {
				r.DrawPixel(i, j, color)
			}
		}
	}
}

// DrawCircle draws a circle in the texture.
//
// It takes the x, y, and radius as parameters.
// No return types.
func (r *Renderer) DrawCircle(x, y, radius int, color int) {
	// create a rectangle to interate over the pixels
	rect := rl.NewRectangle(float32(x-radius), float32(y-radius), float32(radius*2), float32(radius*2))

	// for each pixel in the rectangle
	for i := 0; i < radius*2; i++ {
		for j := 0; j < radius*2; j++ {
			// check if the pixel is in the circle
			if math.Pow(float64(i-x), 2)+math.Pow(float64(j-y), 2) <= math.Pow(float64(radius), 2) {
				r.DrawPixel(int(rect.X)+i, int(rect.Y)+j, color)
			}
		}
	}
}

// DrawCircleLines draws a circle in the texture.
//
// It takes the x, y, and radius as parameters.
// No return types.
func (r *Renderer) DrawCircleLines(x, y, radius int, color int) {
	// calculate the number of points to draw
	points := int(math.Ceil(math.Sqrt(float64(radius * radius * 4))))

	// draw a circle in the texture
	for i := 0; i < points; i++ {
		r.DrawPixel(x+int(math.Cos(float64(i)*2*math.Pi/float64(points))*float64(radius)), y+int(math.Sin(float64(i)*2*math.Pi/float64(points))*float64(radius)), color)
	}
}

// DrawLine draws a line in the texture.
//
// It takes the x1, y1, x2, y2, and color as parameters.
// No return types.
func (r *Renderer) DrawLine(x1, y1, x2, y2 int, color int) {
	// draw a line in the texture using the bresenham algorithm
	dx := x2 - x1
	dy := y2 - y1
	x := x1
	y := y1
	sx, sy := 1, 1
	if dx < 0 {
		sx = -1
		dx = -dx
	}
	if dy < 0 {
		sy = -1
		dy = -dy
	}
	err := dx - dy

	for {
		r.DrawPixel(x, y, color)
		if x == x2 && y == y2 {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x += sx
		}
		if e2 < dx {
			err += dx
			y += sy
		}
	}
}

// DrawPixel draws a pixel in the texture.
//
// It takes the x, y, and color as parameters.
// No return types.
func (r *Renderer) DrawPixel(x, y int, color int) {
	// draw a pixel in the texture
	rl.DrawPixel(int32(x), int32(y), PalletIndexToColor(color))
}

// DrawTriangle draws a triangle in the texture.
//
// It takes the x1, y1, x2, y2, x3, y3, and color as parameters.
// No return types.
func (r *Renderer) DrawTriangle(x1, y1, x2, y2, x3, y3 int, color int) {
	// Sort the three points from top to bottom
	if y1 > y2 {
		x1, x2, y1, y2 = x2, x1, y2, y1
	}
	if y1 > y3 {
		x1, x3, y1, y3 = x3, x1, y3, y1
	}
	if y2 > y3 {
		x2, x3, y2, y3 = x3, x2, y3, y2
	}

	// Calculate the slopes of the two sides of the triangle
	slope1 := float64(x2-x1) / float64(y2-y1)
	slope2 := float64(x3-x1) / float64(y3-y1)

	// Initialize the starting and ending points for each scanline
	xStart := x1
	xEnd := x1

	// Fill in the pixels for each scanline
	for y := y1; y <= y3; y++ {
		for x := xStart; x <= xEnd; x++ {
			r.DrawPixel(x, y, color) // Set the pixel color
		}
		xStart += int(slope1)
		xEnd += int(slope2)
	}
}

// DrawTriangle draws a triangle in the texture.
//
// It takes the x1, y1, x2, y2, x3, y3, and color as parameters.
// No return types.
func (r *Renderer) DrawTriangleLines(x1, y1, x2, y2, x3, y3 int, color int) {
	// draw the lines that make up the triangle
	r.DrawLine(x1, y1, x2, y2, color)
	r.DrawLine(x2, y2, x3, y3, color)
	r.DrawLine(x3, y3, x1, y1, color)
}
