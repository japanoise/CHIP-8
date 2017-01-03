package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

var (
	Screen *sdl.Texture
)

/// InitScreen creates the render target for the CHIP-8 video memory.
///
func InitScreen() {
	var err error

	// create a render target for the display
	Screen, err = Renderer.CreateTexture(sdl.PIXELFORMAT_RGB888, sdl.TEXTUREACCESS_TARGET, 64, 32)
	if err != nil {
		panic(err)
	}
}

/// RefreshScreen with the CHIP-8 video memory.
///
func RefreshScreen() {
	if err := Renderer.SetRenderTarget(Screen); err != nil {
		panic(err)
	}

	// the background color for the screen
	Renderer.SetDrawColor(143, 145, 133, 255)
	Renderer.Clear()

	// set the pixel color
	Renderer.SetDrawColor(17, 29, 43, 255)

	// draw all the pixels
	for p := 0;p < 64 * 32;p++ {
		if VM.Video[p>>3] & (0x80 >> uint(p&7)) != 0 {
			x := int(p&63)
			y := int(p>>6)

			// render the pixel to the screen
			Renderer.DrawPoint(x, y)
		}
	}

	// restore the render target
	Renderer.SetRenderTarget(nil)
}

/// CopyScreen to the render target.
///
func CopyScreen(x, y, scale int32) {
	Renderer.Copy(Screen, nil, &sdl.Rect{
		X: x,
		Y: y,
		W: 64 * scale,
		H: 32 * scale,
	})
}

/// SaveScreen writes a BMP to disk of the current screen.
///
func SaveScreen() {
	// TODO:
}