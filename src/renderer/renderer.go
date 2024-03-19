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
// It takes the game width, game height, and title as parameters and returns a pointer to Renderer.
func NewRenderer(gameWidth, gameHeight int, title string) *Renderer {

	rl.InitWindow(screenWidth, screenHeight, title)
	rl.SetTargetFPS(60)

	target := rl.LoadRenderTexture(int32(gameWidth), int32(gameHeight))
	rl.SetTextureFilter(target.Texture, rl.FilterPoint)

	return &Renderer{
		Target: &target,
	}
}

// DrawOnScreen draws the renderer on the screen.
//
// No parameters.
// No return types.
func (r *Renderer) DrawOnScreen() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	rl.BeginTextureMode(*r.Target)
	rl.ClearBackground(rl.Black)
	gameWidth := float32(r.Target.Texture.Width)
	gameHeight := float32(r.Target.Texture.Height)
	rl.DrawTexturePro(r.Target.Texture,
		rl.NewRectangle(0, 0, gameWidth, -gameHeight),
		rl.NewRectangle(0, 0, float32(screenWidth), float32(screenHeight)),
		rl.NewVector2(0, 0), 0, rl.White)
	rl.EndDrawing()
}

// Close releases the render texture and closes the window.
//
// No parameters.
// No return types.
func (r *Renderer) Close() {
	rl.UnloadRenderTexture(*r.Target)
	rl.CloseWindow()
}

// methods to draw stuff like text or shapes will be her down the line
