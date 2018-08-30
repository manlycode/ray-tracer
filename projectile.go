package projectile

import (
	"github.com/manlycode/ray-tracer/tuple"
)

//Projectile represents a projectile in cartesian space
type Projectile struct {
	Position Tuple
	Velocity Tuple
}

type World struct {
	Gravity Tuple
	Wind    Tuple
}

func Tick(world World, p Projectile) Projectile {
	newPos := tuple.Add(p.Position, p.Velocity)
	newVel := tuple.add(tuple.add(p.Velocity, world.Gravity), world.Wind)
	return Projectile{Position: position, Velocity: velocity}
}
