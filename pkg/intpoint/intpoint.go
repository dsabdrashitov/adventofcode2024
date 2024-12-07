package intpoint

type Point struct {
	X int
	Y int
}

var DIR4 = []Point{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

var DIR8 = []Point{
	{-1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
	{0, -1},
	{-1, -1},
}

func New(x int, y int) Point {
	return Point{x, y}
}

func (t Point) Add(p Point) Point {
	return Point{t.X + p.X, t.Y + p.Y}
}

func (t Point) Sub(p Point) Point {
	return Point{t.X - p.X, t.Y - p.Y}
}

func (t Point) Mult(m int) Point {
	return Point{t.X * m, t.Y * m}
}

func (t Point) InsideStrings(a []string) bool {
	return PointInsideStrings(t, a)
}

func PointInside[T any](p Point, a [][]T) bool {
	if p.X < 0 {
		return false
	}
	if p.X >= len(a) {
		return false
	}
	if p.Y < 0 {
		return false
	}
	if p.Y >= len(a[p.X]) {
		return false
	}
	return true
}

func PointInsideStrings(p Point, a []string) bool {
	if p.X < 0 {
		return false
	}
	if p.X >= len(a) {
		return false
	}
	if p.Y < 0 {
		return false
	}
	if p.Y >= len(a[p.X]) {
		return false
	}
	return true
}
