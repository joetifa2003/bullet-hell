package collision

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/joetifa2003/bullet-hell/pkg/vector"
)

type Line struct {
	Start  vector.Vector
	End    vector.Vector
	Normal vector.Vector
}

func rectToLines(rect rl.Rectangle) [4]Line {
	return [4]Line{
		{ // left
			Start:  vector.New(rect.X, rect.Y+rect.Height),
			End:    vector.New(rect.X, rect.Y),
			Normal: vector.New(-1, 0),
		},
		{ // right
			Start:  vector.New(rect.X+rect.Width, rect.Y),
			End:    vector.New(rect.X+rect.Width, rect.Y+rect.Height),
			Normal: vector.New(1, 0),
		},

		{ // top
			Start:  vector.New(rect.X, rect.Y),
			End:    vector.New(rect.X+rect.Width, rect.Y),
			Normal: vector.New(0, -1),
		},
		{ // bottom
			Start:  vector.New(rect.X+rect.Width, rect.Y+rect.Height),
			End:    vector.New(rect.X, rect.Y+rect.Height),
			Normal: vector.New(0, 1),
		},
	}
}

func MoveCollide(rect rl.Rectangle, dp vector.Vector, otherRects ...rl.Rectangle) (vector.Vector, bool) {
	collided := false
	rectPos := vector.New(rect.X, rect.Y)

	for _, otherRect := range otherRects {
		otherRect.Width += rect.Width
		otherRect.Height += rect.Height
		otherRect.X -= rect.Width
		otherRect.Y -= rect.Height

		movLine := Line{Start: rectPos, End: rectPos.Add(dp)}
		normalVector := vector.Zero()
		minT := float32(1) // min collision time
		for _, line := range rectToLines(otherRect) {
			t, intersects := lineIntersect(line, movLine)
			if intersects {
				if t < minT {
					minT = t
					normalVector = line.Normal
					collided = true
				}
			}
		}

		// if the displacement vector is opposite to the normal vector
		// scale the displacement vector with collision time to prevent it
		// from pentrating into otherRect
		// removing this check makes the entity stuck after collision
		if dp.Dot(normalVector) < 0 {
			scaledDp := dp.Scale(minT)

			// Calculate the remaining displacement vector
			remainingDp := dp.Scale(1 - minT)

			// Project the remaining displacement onto the tangent plane
			tangent := vector.New(normalVector.Y(), normalVector.X())
			tangentDp := remainingDp.Project(tangent)

			// Update the displacement vector to include the tangent displacement
			dp = scaledDp.Add(tangentDp)
		}
	}

	return dp, collided
}

func lineIntersect(l1, l2 Line) (float32, bool) {
	det := (l1.End.X()-l1.Start.X())*(l2.End.Y()-l2.Start.Y()) - (l1.End.Y()-l1.Start.Y())*(l2.End.X()-l2.Start.X())

	if det == 0 {
		return 0, false
	}

	t := ((l2.Start.X()-l1.Start.X())*(l2.End.Y()-l2.Start.Y()) - (l2.Start.Y()-l1.Start.Y())*(l2.End.X()-l2.Start.X())) / det
	u := -((l1.Start.X()-l1.End.X())*(l1.Start.Y()-l2.Start.Y()) - (l1.Start.Y()-l1.End.Y())*(l1.Start.X()-l2.Start.X())) / det

	return u, t >= 0 && t <= 1 && u >= 0 && u <= 1
}
