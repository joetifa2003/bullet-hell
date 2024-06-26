package game

import (
	"bytes"
	_ "embed"
	"image/png"

	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/joetifa2003/bullet-hell/pkg/vector"
)

//go:embed raylib_logo.png
var logo []byte

type TextureImage struct {
	img rl.Image
	tex rl.Texture2D
}

type Player struct {
	BaseEntity
	texture    rl.Texture2D
	pos        vector.Vector
	vel        vector.Vector
	size       vector.Vector
	shootTimer Timer
}

func NewPlayer(pos vector.Vector) *Player {
	return &Player{
		pos:        pos,
		size:       vector.New(50, 50),
		shootTimer: NewTimer(0.2),
	}
}

func (p *Player) Load() {
	img, err := png.Decode(bytes.NewBuffer(logo))
	if err != nil {
		panic(err)
	}
	_ = img
}

func (p *Player) Update(dt float32) {
	p.shootTimer.Update(dt)

	p.vel = p.vel.SetX(0).SetY(0)

	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		p.vel = p.vel.SetY(-1)
	}

	if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		p.vel = p.vel.SetY(1)
	}

	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		p.vel = p.vel.SetX(-1)
	}

	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		p.vel = p.vel.SetX(1)
	}

	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		if p.shootTimer.Done() {
			center := p.pos.Add(p.size.Scale(0.5))
			dir := vector.MouseWorld(p.game.camera).Sub(center).Norm()
			pos := center.Add(dir.Scale(50))
			p.game.AddEntity(&Bullet{
				pos: pos,
				vel: dir,
			})

			p.shootTimer.Reset()
		}
	}

	p.vel = p.vel.Norm().Scale(300 * dt)

	p.vel, _ = p.game.MoveCollide(p, p.vel)

	p.pos = p.pos.Add(p.vel)
	p.game.camera.Target = p.pos.RL()
}

func (p *Player) Draw() {
	rl.DrawRectangle(int32(p.pos.X()), int32(p.pos.Y()), int32(p.size.X()), int32(p.size.Y()), rl.Yellow)
}

func (p *Player) DrawScreen() {
	texture := p.game.GetRenderTexture(Layer1)
	texture.BlendMode = rl.BlendAdditive
	pos := p.game.WorldToScreen(p.pos).Add(vector.New(25, -25))

	rl.BeginTextureMode(texture.RenderTexture2D)
	rl.ClearBackground(rl.Blank)
	rl.DrawCircleGradient(int32(pos.X()), int32(pos.Y()), 50, rl.Fade(rl.White, 0.7), rl.Fade(rl.White, 0.4))
	rl.EndTextureMode()
}

func (p *Player) Pos() vector.Vector { return p.pos }

func (p *Player) CollisionShape() rl.Rectangle {
	return rl.Rectangle{
		X:      p.pos.X(),
		Y:      p.pos.Y(),
		Width:  p.size.X(),
		Height: p.size.Y(),
	}
}
