package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/joetifa2003/bullet-hell/pkg/vector"
)

type Block struct {
	BaseEntity
	pos  vector.Vector
	size vector.Vector
}

func (b *Block) Draw() {
	rl.DrawRectangle(int32(b.pos.X()), int32(b.pos.Y()), int32(b.size.X()), int32(b.size.Y()), rl.Red)
}

func (b *Block) CollisionShape() rl.Rectangle {
	return rl.Rectangle{
		X:      b.pos.X(),
		Y:      b.pos.Y(),
		Width:  b.size.X(),
		Height: b.size.Y(),
	}
}

func NewBlock(pos vector.Vector, size vector.Vector) *Block {
	return &Block{
		pos:  pos,
		size: size,
	}
}
