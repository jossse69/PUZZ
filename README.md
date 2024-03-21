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
 "github.com/jossse69/PUZZ/core"          // import the core package
 "github.com/jossse69/PUZZ/renderer/font" // import the font package
)

var (
 game *core.Game // global game variable

 // a simple font variable for the text
 testFont *font.Font
)

func main() {
 // low res of the screen of 320x240 for retro look. (this a recommended setting, but you can change it if you want)
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

 // load the font
 // note, you need to use a Codepage 437 font for this, you can find good ones in the dwarf fortress wiki's tileset repository: https://dwarffortresswiki.org/Tileset_repository
 // making your own font is also possible if you want!
 // the font's background is transparent, so if you are editing or drawing a font, you need to use a transparent background
 testFont = font.LoadFontFromImage("Zaratustra_msx_transparent.png", 8, 8, 16, rune(' '))
}

func update(dt float32) {
 // Update game logic here, use dt for time since last frame (useful for frame rate independent movement)
}

func draw() {
 // draw any elements here
 // here is an example of drawing a string of text

 // first clear the screen
 game.Renderer.Fill(0)

 // draw the said text
 game.Renderer.DrawText(testFont, 0, 0, "Hello, World!")
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
