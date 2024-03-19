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

import "github.com/jossse69/PUZZ" // import the engine

const (
    // width and height of the game
    WIDTH = 320
    HEIGHT = 200
)

func main() {
    r := puzz.Render.NewRenderer(WIDTH, HEIGHT, "My super duper epic game title") // create a new renderer so we can draw stuff

    // main loop
    for {

        r.Clear() // clear the screen for the next frame

        // draw stuff here
        // example: Draw a string of text
        r.Print(0, 0, "Hello World!")

        r.DrawOnScreen() // draw on the screen
    }

    // don't forget to close the renderer!
    r.Close()
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
