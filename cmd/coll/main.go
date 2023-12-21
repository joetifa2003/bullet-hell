package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/joetifa2003/bullet-hell/cmd/pkg/collision"
)

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rect := rl.Rectangle{
		X:      200,
		Y:      100,
		Width:  100,
		Height: 300,
	}

	startingPos := rl.Vector2{X: 0, Y: 0}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawFPS(10, 10)

		endPos := rl.Vector2{X: float32(rl.GetMouseX()), Y: float32(rl.GetMouseY())}
		startRec := rl.Rectangle{X: startingPos.X, Y: startingPos.Y, Width: 25, Height: 25}

		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			startingPos = rl.Vector2{X: float32(rl.GetMouseX()), Y: float32(rl.GetMouseY())}
		}

		newRect := collision.MoveCollide(startRec, rl.Vector2Subtract(endPos, startingPos), rect)
		rl.DrawRectangleRec(rl.Rectangle{X: startRec.X + newRect.X, Y: startRec.Y + newRect.Y, Width: 25, Height: 25}, rl.Green)

		rl.DrawRectangleRec(startRec, rl.Red)
		rl.DrawRectangleRec(rect, rl.Fade(rl.Red, 0.5))
		// rl.DrawLine(int32(startingPos.X), int32(startingPos.Y), int32(endPos.X), int32(endPos.Y), rl.Black)

		rl.EndDrawing()
	}
}
