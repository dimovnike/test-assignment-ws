package main

// Obstacle is an obstacle on the grid
// all squares with x1 <= x <= x2 and y1 <= y <= y2 are occupied.
type Obstacle struct {
	X1, X2 int
	Y1, Y2 int
}

// Grid defines a grid with obstacles
type Grid struct {
	Width     int
	Height    int
	Obstacles []Obstacle
}

type Point struct {
	X, Y int
}

// Hop returns a destination point after Point hops at given speed.
func (p Point) Hop(s speed) Point {
	return Point{p.X + s.x, p.Y + s.y}
}

// MinHops calculates the optimal number of hops
func (g *Grid) MinHops(start Point, end Point) int {
	// Using the BFS algorithm due to its property to find shortest path, which in our case is the number of hops.
	// The algorithm is modified and visits cells multiple times because it can arrive
	// at different speeds at them and so can reach different new cells.

	type pointSpeed struct {
		Point
		speed speed
	}

	type pointHops struct {
		pointSpeed
		minHops int
	}

	// using pointSpeed type to track visited cells as per above.
	visited := make(map[pointSpeed]bool, 0)

	// creating a FIFO queue
	var q Queue[pointHops]

	q.Queue(pointHops{pointSpeed{start, speed{}}, 0})

	minHops := -1

	for {
		ps, ok := q.Deqeue()
		if !ok {
			break
		}

		if ps.Point == end {
			minHops = ps.minHops
			break
		}

		for _, s := range GetSpeeds(ps.speed) {
			hopPoint := ps.Point.Hop(s)
			hopPointSpeed := pointSpeed{hopPoint, s}

			if visited[hopPointSpeed] || !g.isValid(hopPoint) || g.isObstacle(hopPoint) {
				continue
			}

			visited[hopPointSpeed] = true

			q.Queue(pointHops{hopPointSpeed, ps.minHops + 1})
		}
	}

	return minHops
}

// Print prints a grid with its obstacles, this is for testing purposes.
func (g *Grid) Print() string {
	result := make([]byte, 0, g.Height*g.Width)

	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			pointStr := "."
			if g.isObstacle(Point{x, y}) {
				pointStr = "X"
			}

			result = append(result, pointStr...)
		}

		result = append(result, "\n"...)
	}

	return string(result)
}

// isValid checks if the point is a valid point on the grid (i.e. is whithin the grid bounds)
func (g *Grid) isValid(p Point) bool {
	return p.X >= 0 && p.X < g.Width && p.Y >= 0 && p.Y < g.Height
}

// isObstacle checks if the point is on an obstacle
func (g *Grid) isObstacle(p Point) bool {
	for _, o := range g.Obstacles {
		if p.X >= o.X1 && p.X <= o.X2 && p.Y >= o.Y1 && p.Y <= o.Y2 {
			return true
		}
	}

	return false
}
