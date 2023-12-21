package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/joetifa2003/bullet-hell/pkg/collision"
	"github.com/joetifa2003/bullet-hell/pkg/vector"
)

type Entity interface {
	Update(dt float32)
	Draw()
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
func (b *BaseEntity) Draw()             {}
func (b *BaseEntity) SetIndex(idx int)  { b.index = idx }
func (b *BaseEntity) Index() int        { return b.index }
func (b *BaseEntity) SetGame(g *Game)   { b.game = g }

type Game struct {
	camera           *rl.Camera2D
	entities         []Entity
	entitiesToRemove []Entity
	entitiesToAdd    []Entity
}

func (g *Game) AddEntity(e Entity) {
	g.entitiesToAdd = append(g.entitiesToAdd, e)
}

func (g *Game) RemoveEntity(e Entity) {
	g.entitiesToRemove = append(g.entitiesToRemove, e)
}

func (g *Game) Update(dt float32) {
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

func (g *Game) MoveCollide(collidable Collidable, vel vector.Vector) (vector.Vector, bool) {
	collided := false
	for _, e := range g.entities {
		if e, ok := e.(Collidable); ok {
			var c bool
			vel, c = collision.MoveCollide(collidable.CollisionShape(), vel, e.CollisionShape())
			if c && !collided {
				collided = true
			}
		}
	}

	return vel, collided
}

func NewGame(camera *rl.Camera2D) *Game {
	return &Game{
		camera: camera,
	}
}
