package renderer

import (
	rl "github.com/gen2brain/raylib-go/raylib"
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
	rl.DrawTexturePro(r.Target.Texture,
		rl.NewRectangle(0, 0, gameWidth, -gameHeight),
		rl.NewRectangle((float32(screenWidth)/2)-(gameWidth/2), (float32(screenHeight)/2)-(gameHeight/2), float32(screenWidth), float32(screenHeight)),
		rl.NewVector2(0, 0), 0, rl.White)
	rl.EndDrawing() // End drawing to the main window
}

// PalletIndexToColor converts a pallet index to a color.
//
// It takes the index of the color in the pallet.
// Returns the corresponding color.
func PalletIndexToColor(index int) rl.Color {

	// taken form the anb16 palette (https://lospec.com/palette-list/anb16)

	// the list of colors
	colors := []int{
		0x0a080d,
		0x697594,
		0xdfe9f5,
		0xf7aaa8,
		0xd4689a,
		0x782c96,
		0xe83562,
		0xf2825c,
		0xffc76e,
		0x88c44d,
		0x3f9e59,
		0x373461,
		0x4854a8,
		0x7199d9,
		0x9e5252,
		0x4d2536,
	}

	// create a color from the index
	selectedColor := colors[index]
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

// Fill fills the screen with a color.
//
// It takes the color as a parameter.
// No return types.
func (r *Renderer) Fill(color int) {
	rl.ClearBackground(PalletIndexToColor(color))
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
	rl.DrawRectangle(int32(x), int32(y), int32(width), int32(height), PalletIndexToColor(color))
}

// DrawRectangleLines draws a rectangle in the texture.
//
// It takes the x, y, width, height, and color as parameters.
// No return types.
func (r *Renderer) DrawRectangleLines(x, y, width, height int, color int) {
	// draw a rectangle in the texture
	rl.DrawRectangleLines(int32(x), int32(y), int32(width), int32(height), PalletIndexToColor(color))
}

// DrawCircle draws a circle in the texture.
//
// It takes the x, y, and radius as parameters.
// No return types.
func (r *Renderer) DrawCircle(x, y, radius int, color int) {
	// draw a circle in the texture
	rl.DrawCircleV(rl.NewVector2(float32(x), float32(y)), float32(radius), PalletIndexToColor(color))
}

// DrawCircleLines draws a circle in the texture.
//
// It takes the x, y, and radius as parameters.
// No return types.
func (r *Renderer) DrawCircleLines(x, y, radius int, color int) {
	// draw a circle in the texture
	rl.DrawCircleLines(int32(x), int32(y), float32(radius), PalletIndexToColor(color))
}

// DrawLine draws a line in the texture.
//
// It takes the x1, y1, x2, y2, and color as parameters.
// No return types.
func (r *Renderer) DrawLine(x1, y1, x2, y2 int, color int) {
	// draw a line in the texture
	rl.DrawLine(int32(x1), int32(y1), int32(x2), int32(y2), PalletIndexToColor(color))
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
	// draw a triangle in the texture
	v1 := rl.NewVector2(float32(x1), float32(y1))
	v2 := rl.NewVector2(float32(x2), float32(y2))
	v3 := rl.NewVector2(float32(x3), float32(y3))
	rl.DrawTriangle(v1, v2, v3, PalletIndexToColor(color))
}

// DrawTriangle draws a triangle in the texture.
//
// It takes the x1, y1, x2, y2, x3, y3, and color as parameters.
// No return types.
func (r *Renderer) DrawTriangleLines(x1, y1, x2, y2, x3, y3 int, color int) {
	// draw a triangle in the texture
	v1 := rl.NewVector2(float32(x1), float32(y1))
	v2 := rl.NewVector2(float32(x2), float32(y2))
	v3 := rl.NewVector2(float32(x3), float32(y3))
	rl.DrawTriangleLines(v1, v2, v3, PalletIndexToColor(color))
}
