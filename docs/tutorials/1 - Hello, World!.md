# Tutorial - Hello, World!
This tutorial will show you how to create a simple game in the [PUZZ Engine](https://github.com/jossse69/PUZZ), and explan how the engine works, and more basic concepts.

# 1 - What is PUZZ?
PUZZ is a game engine for retro looking games. It is made in 2024 by José (called online with the nickname 'tomato'), and it is based on 'Fantasy consoles' (wich are 'console emulators' for consoles that never did exist, adding a retro feel via restrictions such as lower resolution, limited colors, etc), a good example of these consoles is the PICO-8, the project that started this consoles. PUZZ has the same philosophy as a fantasy consoles but it is not a fantasy console, but a retro game engine, as ther is not limit on the amount of 'stuff' you can put in the game (content: graphics, sounds, etc). its 100% made in Golang.

# 2 - What is Go/Golang?
Go (also called as Golang) is a programming language that is used for many wide software development. This includes cloud and network systems, command-line interfaces, web development, and, you guessed it, games and engines. This tutorial will show you how to install it, and how to use it.

# 3 - Settin' up
Our first step is to get Go installed. You can download it from the [official website](https://golang.org/dl/) and install it in your computer. following the instructions, you will be able to use the 'go' command. you can test it via this command: `go version`. this should return the version of Go you have installed. its aslo recomended to learn it as well, so that you can use it in your project. <br/>

Now that you have Go installed, you need to install the [PUZZ Engine](https://github.com/jossse69/PUZZ). you can do this via the command: `go get github.com/jossse69/PUZZ`. also if your IDE/compiler is telling about package errors about missing depandencies, you can istall them via the command: `go mod tidy`.

# 4 - A new game
Lets create a project, you can create a Go project like so:
- first, create a folder for your project
- then, initialize the project: `go mod init <project-name>`
- intall the dependencies as said above.

This will create a `go.mod` file and a `go.sum` file. don't poke into this files too much as of now. anyways, lets create a `main.go` file, this is were our program will run once its started.

So lets create a simple basic base for a go file:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```
this will print "Hello, World!" into the console. you can now run it with the command: `go run main.go` to complie and run it. well, this isnt much of a game, but it will shows if everything is working as it should.

lets use PUZZ in our project. import it in our `main.go` file:
```go
import (
    "github.com/jossse69/PUZZ/core"
    // add any other packages you need
)
```
the PUZZ engine is split into packages. for core game engine, you can find it in the `core` package. now, create a variable of type `core.Game` and assign it to the `game` variable. this will be the main game. you can then use the `game` variable to start the game:
```go
var game *core.Game 

// inside func main()...
game = core.NewGame(220, 140, "Hello, World!") // this will set the game resolution and title
game.Run() // this will start the game loop
```
This will create a new game and run it, running our program, we see it makes a window with the title "Hello, World!", but the window is empty. this is because we have not telled our game to do anything yet, that includes rendering. to do it we can set some 'hooks', this are methods that are run when something happens, the `core.Game` has 3 hooks: `Load`, `Update` and `Draw`. this are the most basic hooks, made for the game loop. so, create 3 new methods:
```go
func Load() {
    // load any assets here
}

func Update() {
    // update the game here
}

func Draw() {
    // draw any elements here
}
```
Back in the main func, add the hooks to the game:
```go
func main() {
    game = core.NewGame(220, 140, "Hello, World!")
    game.Load = Load
    game.Update = Update
    game.Draw = Draw
    game.Run()
}
```
Now that we have hooks, we can inplement them. for this tutorial we will make our game draw sone text. lets frist import the font package:
```go
import (
    "github.com/jossse69/PUZZ/core"
    "github.com/jossse69/PUZZ/renderer/font"
)
```
Before we can draw text, we need to load a font. this is done in the `Load` method using the `font.LoadFontFromImage` method:
```go
func Load() {
    // make sure you set a 'font' variable on the top of this file!
    font = font.LoadFontFromImage("assets/font.png", 8, 8, 16, rune(' '))
}
```
for this tutorial we will use a font that is made in the `assets` folder. so lets create a new file called `assets/font.png` and copy this image into it:
![image](assets/font.png)

When making PUZZ fonts, your fonts must be in the Code Page 437 format. PUZZ already handles this, but your font needs to have a trasnsparent background, and your fonts also need to have a white foreground. you can find good fonts in the dwarf fortress wiki's tileset repository [here](https://dwarffortresswiki.org/Tileset_repository), you will need to edit them for tings like the background and the white foreground, but that should be enough. <br/>

finally, lets create our `Draw` method and draw our text:
```go
func Draw() {
    // first clear the screen
	game.Renderer.Fill(0)

	// draw the said text
	game.Renderer.DrawText(font, 0, 0, "Hello, World!")
}
```
Congrats, you made your frist 'game' in PUZZ. even if its not really a game, but it if yo followed the instructions, you should be able to make a window with the title "Hello, World!" in it and see it.