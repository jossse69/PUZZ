# 👾 PUZZ ENGINE 👾
a Golang engine for small retro games.
create anything you can imagine! a platformer? a puzzle game? heck even a 3D game!? you can do it! :D

## Features
- easy to use and get started.
- extensive docs for all the lil' details of stuff.
- 100% Golang, Go lovers only!
- Pixel based graphics/rendering.
- simple object system to get the job done for you!
- simple 8 bit sounds for sfx and music.
- tilesets and tilemaps, powered by [Tiled](https://www.mapeditor.org/)
- 100% open source, anyone can contribute to this engine, anywere, anytime.
- 100% free to use!
- More to come here!

## Getting Started
first, if you have [Go](https://go.dev/) installed, run this command: </br>
`go get github.com/jossse69/PUZZ` </br>
this will download all the dependencies, you are all set!

## Usage
here is a simple 'Hello World' example:
```go
package main

import (
 "github.com/jossse69/PUZZ/core" // import the core package 
)

var game *core.Game // global game variable

func main() {
 // low res of the screen of 320x240 for retro look. (this a recommended setting)
 game = core.NewGame(220, 140, "Hello, World!")

 // setup the load, update and draw hooks
 game.Load = load
 game.Update = update
 game.Draw = draw

 // finally, run the game
 game.Run()
}

func load() {
 // Load any game resources here
}

func update(dt float32) {
 // Update game logic here, use dt for time since last frame (useful for frame rate independent movement)
}

func draw() {
 // draw any elements here
 // e.g. line for the game rect

 // first clear the screen
 game.Renderer.Fill(0)

 // draw a line from 0,0 to the top right of the screen
 game.Renderer.DrawLine(0, 0, int(game.Renderer.Target.Texture.Width), int(game.Renderer.Target.Texture.Height), 6) // id color of 6 is red

 // then a line from 0,0 to the right of the screen
 game.Renderer.DrawLine(0, 0, int(game.Renderer.Target.Texture.Width), 0, 6)

 // then a line from 0,0 to the top of the screen
 game.Renderer.DrawLine(0, 0, 0, int(game.Renderer.Target.Texture.Height), 6)
}
``` 

want to know more? check out the [docs](https://github.com/jossse69/PUZZ/tree/master/docs) for more information! :D

## License
[MIT](https://github.com/jossse69/PUZZ/blob/master/LICENSE) license is used.

## Contributing
If you want to contribute, check out the [contribution guide](https://github.com/jossse69/PUZZ/blob/master/CONTRIBUTING.md) for more information. all contributions are welcome! :D

## Credits
- [jossse69 (AKA Tomato)](https://github.com/jossse69) - author
- more to come here...
