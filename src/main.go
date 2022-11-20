package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(SCREEN_SIZE*PIXEL_SIZE, SCREEN_SIZE*PIXEL_SIZE, "4BOD")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		DrawBitmap(BMP_SPLASH)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
