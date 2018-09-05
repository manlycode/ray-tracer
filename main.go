package main

import (
	"fmt"

	"github.com/manlycode/ray-tracer/point"
	"github.com/manlycode/ray-tracer/tuple"
	"github.com/manlycode/ray-tracer/vector"
)

//Projectile represents a projectile in cartesian space
type Projectile struct {
	Position point.Point
	Velocity vector.Vector
}

// World represents an origin relative to things in it
type World struct {
	Gravity vector.Vector
	Wind    vector.Vector
}

func tick(world World, p Projectile) Projectile {
	newPos := point.Point{tuple.Add(p.Position, p.Velocity)}
	newVel := vector.Vector{tuple.Add(tuple.Add(p.Velocity, world.Gravity), world.Wind)}
	return Projectile{newPos, newVel}
}

func main() {
	projectile := Projectile{
		point.New(0, 0, 0),
		vector.New(1, 2, 0),
	}

	world := World{
		vector.New(-0.5, -0.5, 0),
		vector.New(0, 0, 0),
	}

	fmt.Println("Hello")

	sum := 1
	for sum < 200 {
		projectile = tick(world, projectile)
		fmt.Println(projectile.Position.GetTuple())
		sum++
	}
}
