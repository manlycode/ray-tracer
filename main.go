package main

import (
	"fmt"

	"github.com/manlycode/ray-tracer/tuple"
)

//Projectile represents a projectile in cartesian space
type Projectile struct {
	Position tuple.Point
	Velocity tuple.Vector
}

// World represents an origin relative to things in it
type World struct {
	Gravity tuple.Vector
	Wind    tuple.Vector
}

func tick(world World, p Projectile) Projectile {
	newPos := tuple.Add(p.Position, p.Velocity)
	newVel := tuple.Add(tuple.Add(p.Velocity, world.Gravity), world.Wind)
	return Projectile{newPos.(tuple.Point), newVel.(tuple.Vector)}
}

func main() {
	projectile := Projectile{
		tuple.NewPoint(0, 0, 0),
		tuple.NewVector(1, 2, 0),
	}

	world := World{
		tuple.NewVector(-0.5, -0.5, 0),
		tuple.NewVector(0, 0, 0),
	}

	fmt.Println("Hello")

	sum := 1
	for sum < 200 {
		projectile = tick(world, projectile)
		fmt.Println(projectile.Position.GetTuple())
		sum++
	}
}
