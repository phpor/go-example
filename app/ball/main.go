package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y, Z float64
}

type BoundingSphere struct {
	Center Point
	Radius float64
}

func calculateBoundingSphere(points []Point) BoundingSphere {
	// Step 1: Find min and max coordinates
	minX, maxX := math.MaxFloat64, math.SmallestNonzeroFloat64
	minY, maxY := math.MaxFloat64, math.SmallestNonzeroFloat64
	minZ, maxZ := math.MaxFloat64, math.SmallestNonzeroFloat64
	for _, point := range points {
		if point.X < minX {
			minX = point.X
		}
		if point.X > maxX {
			maxX = point.X
		}
		if point.Y < minY {
			minY = point.Y
		}
		if point.Y > maxY {
			maxY = point.Y
		}
		if point.Z < minZ {
			minZ = point.Z
		}
		if point.Z > maxZ {
			maxZ = point.Z
		}
	}

	// Step 2: Calculate radius using formula for sphere passing through 3 points (x1, y1, z1), (x2, y2, z2), (x3, y3, z3)
	// Radius = sqrt((x2 - x1)^2 + (y2 - y1)^2 + (z2 - z1)^2 + (x3 - x1)^2 + (y3 - y1)^2 + (z3 - z1)^2) / 2) - sqrt((x3 - x1)^2 + (y3 - y1)^2 + (z3 - z1)^2) / 2)
	// In this case, we use the min and max coordinates as the 3 points.
	radius := math.Sqrt(math.Pow(maxX-minX, 2)+math.Pow(maxY-minY, 2)+math.Pow(maxZ-minZ, 2))/2 - math.Sqrt(math.Pow(maxX-minX, 2)+math.Pow(maxY-minY, 2)+math.Pow(maxZ-minZ, 2))/2

	// Step 3: Calculate center using the min and max coordinates and the radius.
	center := Point{
		X: (minX + maxX) / 2,
		Y: (minY + maxY) / 2,
		Z: (minZ + maxZ) / 2,
	}
	return BoundingSphere{Center: center, Radius: radius}
}

func main() {
	points := []Point{
		{-1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}
	boundingSphere := calculateBoundingSphere(points)
	fmt.Printf("Bounding Sphere Center: (%f, %f, %f)\n", boundingSphere.Center.X, boundingSphere.Center.Y, boundingSphere.Center.Z)
	fmt.Printf("Bounding Sphere Radius: %f\n", boundingSphere.Radius)
}
