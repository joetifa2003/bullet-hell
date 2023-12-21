package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/joetifa2003/bullet-hell/pkg/game"
	"github.com/joetifa2003/bullet-hell/pkg/vector"
)

const DEBUG_COLLISON = false

const (
	GAME_WIDTH  = 1280
	GAME_HEIGHT = 720
)

func main() {
	rl.InitWindow(GAME_WIDTH, GAME_HEIGHT, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	// rl.SetTargetFPS(60)

	camera := rl.NewCamera2D(rl.NewVector2(GAME_WIDTH/2, GAME_HEIGHT/2), rl.NewVector2(0, 0), 0, 1)

	g := game.NewGame(&camera)

	player := game.NewPlayer(vector.New(100, 100))
	g.AddEntity(player)
	g.AddEntity(game.NewBlock(vector.New(0, 0), vector.New(800, 50)))
	g.AddEntity(game.NewBlock(vector.Zero(), vector.New(50, 800)))

	// lightTexture := rl.LoadRenderTexture(GAME_WIDTH, GAME_HEIGHT)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		dt := rl.GetFrameTime()

		g.Update(dt)

		rl.BeginMode2D(camera)
		g.Draw()
		rl.EndMode2D()

		// rl.BeginTextureMode(lightTexture)
		// rl.ClearBackground(rl.Fade(rl.Black, 0.7))
		// circleCenter := vector.New(player.pos.X()+25, player.pos.Y()-25).RL()
		// circleCenter = rl.GetWorldToScreen2D(circleCenter, camera)
		// rl.DrawCircleV(circleCenter, 50, rl.Fade(rl.White, 0.5))
		// rl.EndTextureMode()
		//
		// rl.BeginBlendMode(rl.BlendMultiplied)
		// rl.DrawTexture(lightTexture.Texture, 0, 0, rl.White)
		// rl.EndBlendMode()

		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
}
