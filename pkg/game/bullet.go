package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/joetifa2003/bullet-hell/pkg/vector"
)

type Bullet struct {
	BaseEntity
	pos vector.Vector
	vel vector.Vector
}

func NewBullet(dir vector.Vector) Bullet {
	return Bullet{
		vel: dir,
	}
}

func (b *Bullet) Update(dt float32) {
	b.vel = b.vel.Norm().Scale(300 * dt)
	var collided bool
	b.vel, collided = b.game.MoveCollide(b, b.vel)
	if collided {
		b.game.RemoveEntity(b)
	}
	b.pos = b.pos.Add(b.vel)
}

func (b *Bullet) Draw() {
	rl.DrawRectangle(int32(b.pos.X()), int32(b.pos.Y()), 5, 5, rl.Blue)
}

func (b *Bullet) CollisionShape() rl.Rectangle {
	return rl.Rectangle{
		X:      b.pos.X(),
		Y:      b.pos.Y(),
		Width:  5,
		Height: 5,
	}
}
