package main

import (
	"fmt"

	"github.com/manlycode/ray-tracer/canvas"
	"github.com/manlycode/ray-tracer/color"
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
	canvas := canvas.New(500, 500)

	projectile := Projectile{
		point.New(0, 0, 0),
		vector.New(5, 20, 0),
	}

	world := World{
		vector.New(0, -0.5, 0),
		vector.New(0, 0, 0),
	}

	sum := 1
	for sum < 200 {
		posTuple := projectile.Position.GetTuple()
		fmt.Println(posTuple)
		canvas.WritePixelFlipped(int(projectile.Position.X()), int(projectile.Position.Y()), color.Red())

		projectile = tick(world, projectile)
		sum++
	}
	canvas.Save("projectile.ppm")
}
