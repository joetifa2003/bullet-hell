package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/joetifa2003/bullet-hell/pkg/collision"
	"github.com/joetifa2003/bullet-hell/pkg/vector"
)

const (
	GAME_WIDTH  = 1280
	GAME_HEIGHT = 720
)

type ScreenLayer int

const (
	Layer1 ScreenLayer = iota
	LayerEnd
)

type Entity interface {
	Update(dt float32)
	Draw()
	DrawScreen()
	Load()

	SetIndex(idx int)
	Index() int
	SetGame(g *Game)
}

type Collidable interface {
	CollisionShape() rl.Rectangle
}

type BaseEntity struct {
	game             *Game
	markedForRemoval bool
	index            int
}

func (b *BaseEntity) Update(dt float32) {}

func (b *BaseEntity) Draw() {}

func (b *BaseEntity) DrawScreen() {}

func (b *BaseEntity) SetIndex(idx int) { b.index = idx }

func (b *BaseEntity) Index() int { return b.index }

func (b *BaseEntity) SetGame(g *Game) { b.game = g }

func (b *BaseEntity) Load() {}

type ScreenTexture struct {
	rl.RenderTexture2D
	BlendMode rl.BlendMode
}

type Game struct {
	camera           *rl.Camera2D
	entities         []Entity
	entitiesToRemove []Entity
	entitiesToAdd    []Entity

	screenTextures []*ScreenTexture
}

func (g *Game) AddEntity(e Entity) {
	g.entitiesToAdd = append(g.entitiesToAdd, e)
}

func (g *Game) RemoveEntity(e Entity) {
	g.entitiesToRemove = append(g.entitiesToRemove, e)
}

func (g *Game) Update(dt float32) {
	if rl.IsKeyPressed(rl.KeyF11) {
		rl.ToggleFullscreen()
	}

	for _, e := range g.entities {
		e.Update(dt)
	}

	for _, e := range g.entitiesToRemove {
		lastEntity := g.entities[len(g.entities)-1]
		lastEntity.SetIndex(e.Index())
		g.entities[e.Index()] = lastEntity
		g.entities = g.entities[:len(g.entities)-1]
	}

	for i, e := range g.entitiesToAdd {
		e.Load()
		e.SetGame(g)
		e.SetIndex(len(g.entities) + i)
	}

	g.entities = append(g.entities, g.entitiesToAdd...)

	g.entitiesToRemove = g.entitiesToRemove[:0]
	g.entitiesToAdd = g.entitiesToAdd[:0]
}

func (g *Game) Draw() {
	for _, e := range g.entities {
		e.Draw()
	}
}

func (g *Game) DrawScreen() {
	for _, e := range g.entities {
		e.DrawScreen()
	}
}

func (g *Game) Loop() {
	rl.InitWindow(GAME_WIDTH, GAME_HEIGHT, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	// rl.SetTargetFPS(60)

	player := NewPlayer(vector.New(100, 100))
	g.AddEntity(player)
	g.AddEntity(NewBlock(vector.New(0, 0), vector.New(800, 50)))
	g.AddEntity(NewBlock(vector.Zero(), vector.New(50, 800)))

	for i := 0; i < int(LayerEnd); i++ {
		g.screenTextures = append(g.screenTextures, &ScreenTexture{RenderTexture2D: rl.LoadRenderTexture(GAME_WIDTH, GAME_HEIGHT)})
	}

	darkenTexture := rl.LoadRenderTexture(GAME_WIDTH, GAME_HEIGHT)
	defer rl.UnloadRenderTexture(darkenTexture)

	rl.BeginTextureMode(darkenTexture)
	rl.ClearBackground(rl.Fade(rl.Black, 0.34))
	rl.EndTextureMode()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		dt := rl.GetFrameTime()

		g.Update(dt)

		rl.BeginMode2D(*g.camera)
		g.Draw()
		rl.EndMode2D()

		g.DrawScreen()

		rl.DrawTexture(darkenTexture.Texture, 0, 0, rl.White)
		for _, t := range g.screenTextures {
			rl.BeginBlendMode(t.BlendMode)
			rl.DrawTexture(t.Texture, 0, 0, rl.White)
			rl.EndBlendMode()
		}

		// rl.BeginTextureMode(lightTexture)
		// rl.ClearBackground(rl.Fade(rl.Black, 0.4))
		//
		// circleCenter := vector.New(player.Pos().X()+25, player.Pos().Y()-25).RL()
		// circleCenter = rl.GetWorldToScreen2D(circleCenter, *g.camera)
		// rl.DrawCircleGradient(int32(circleCenter.X), int32(circleCenter.Y), 50, rl.Fade(rl.White, 0.5), rl.Fade(rl.White, 0.8))
		// rl.EndTextureMode()
		//
		// rl.BeginBlendMode(rl.BlendMultiplied)
		// rl.DrawTexture(lightTexture.Texture, 0, 0, rl.White)
		// rl.EndBlendMode()

		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
}

func (g *Game) MoveCollide(collidable Collidable, vel vector.Vector) (vector.Vector, bool) {
	collided := false
	for _, e := range g.entities {
		if e, ok := e.(Collidable); ok {
			var c bool
			vel, c = collision.MoveCollide(collidable.CollisionShape(), vel, e.CollisionShape())
			if c {
				collided = true
			}
		}
	}

	return vel, collided
}

func (g *Game) GetRenderTexture(layer ScreenLayer) *ScreenTexture {
	return g.screenTextures[layer]
}

func (g *Game) WorldToScreen(pos vector.Vector) vector.Vector {
	return vector.FromRL(rl.GetWorldToScreen2D(pos.RL(), *g.camera))
}

func NewGame() *Game {
	camera := rl.NewCamera2D(rl.NewVector2(GAME_WIDTH/2, GAME_HEIGHT/2), rl.NewVector2(0, 0), 0, 1)
	g := &Game{
		camera: &camera,
	}

	return g
}
