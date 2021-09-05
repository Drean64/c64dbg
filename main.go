package main

import (
	"github.com/Drean64/c64"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	c := c64.Make(c64.NTSC)
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
}
