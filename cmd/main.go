package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	windowWidth  int32 = 800
	windowHeight int32 = 600
	window       *sdl.Window
	// renderer     *sdl.Renderer
	isRunning bool
)

func main() {
	var err error

	isRunning, err = initWindow()
	if err != nil {
		fmt.Printf("error initing window: %v", err)
		os.Exit(1)
	}

	err = setup()
	if err != nil {
		fmt.Printf("error in setup: %v", err)
		os.Exit(1)
	}

	for isRunning {
		processInput()
		update()
		render()
	}

	cleanup()
}

func initWindow() (bool, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return false, fmt.Errorf("error initialising SDL: %v", err)
	}

	displayMode, err := sdl.GetCurrentDisplayMode(0)
	if err != nil {
		return false, fmt.Errorf("error getting display mode: %v", err)
	}
	windowWidth = displayMode.W
	windowHeight = displayMode.H

	window, err = sdl.CreateWindow(
		"",
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		windowWidth,
		windowHeight,
		sdl.WINDOW_BORDERLESS,
	)
	if err != nil {
		return false, fmt.Errorf("error creating SDL window: %v", err)
	}

	// renderer, err = sdl.CreateRenderer(window, -1, 0)
	// if err != nil {
	// 	return false, fmt.Errorf("error creating SDL renderer: %v", err)
	// }

	return true, nil
}

func setup() error {

	return nil
}

func processInput() {
	event := sdl.PollEvent()

	switch t := event.(type) {
	case *sdl.QuitEvent:
		isRunning = false
	case *sdl.KeyboardEvent:
		if t.Keysym.Sym == sdl.K_ESCAPE {
			isRunning = false
		}
	}
}

func update() {

}

func render() {
	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	_ = surface.FillRect(nil, 0)

	rect := sdl.Rect{0, 0, 200, 200}
	_ = surface.FillRect(&rect, 0xffff0000)
	_ = window.UpdateSurface()
}

func cleanup() {
	// renderer.Destroy()
	_ = window.Destroy()
	sdl.Quit()
}
