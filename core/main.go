package core

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jossse69/PUZZ/ecs"
	"github.com/jossse69/PUZZ/renderer"
)

// Game represents the core engine structure.
type Game struct {
	Renderer   *renderer.Renderer
	Systems    *ecs.SystemManager
	Entities   *ecs.EntityManager
	Components *ecs.ComponentManager

	// Load calls once when the game starts, useful for loading resources.
	Load func()

	// Update is called every frame.
	// It receives the current frame time in seconds. (AKA 'delta time')
	Update func(dt float32)

	// Draw is called every frame, but after Update and after drawing is started, so it can be used for drawing game graphics.
	Draw func()
}

// NewGame initializes a new PUZZ game instance.
func NewGame(gameWidth, gameHeight int, title string) *Game {
	// Initialize the Raylib window and other settings here if not already done in Renderer.
	game := &Game{
		Renderer:   renderer.NewRenderer(gameWidth, gameHeight, title),
		Systems:    ecs.NewSystemManager(),
		Entities:   ecs.NewEntityManager(),
		Components: ecs.NewComponentManager(),
	}

	return game
}

// Run starts the main game loop.
func (game *Game) Run() {
	if game.Load != nil {
		game.Load()
	}
	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()

		// Check if the Update function is defined before calling it
		if game.Update == nil {
			panic("Update function is not defined! Please define it before running the game.")
		}
		game.Update(dt)

		game.Renderer.StartDrawing()

		// Check if the Draw function is defined before attempting to call it
		if game.Draw == nil {
			panic("Draw function is not defined! Please define it before running the game.")
		}
		game.Draw()

		game.Renderer.DrawOnScreen()
	}

	game.Close()
}

// Close cleans up resources and closes the game window.
func (game *Game) Close() {
	game.Renderer.Close()
}

// Additional methods to manage game state, handle input, etc., can be added here.
