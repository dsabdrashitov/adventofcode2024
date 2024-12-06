package main

type Point struct {
	x int
	y int
}

var DIR4 = []Point{
	Point{-1, 0},
	Point{0, 1},
	Point{1, 0},
	Point{0, -1},
}

var DIR8 = []Point{
	Point{-1, 0},
	Point{-1, 1},
	Point{0, 1},
	Point{1, 1},
	Point{1, 0},
	Point{1, -1},
	Point{0, -1},
	Point{-1, -1},
}

func (t Point) Add(p Point) Point {
	return Point{t.x + p.x, t.y + p.y}
}

func (t Point) Sub(p Point) Point {
	return Point{t.x - p.x, t.y - p.y}
}

func (t Point) Mult(m int) Point {
	return Point{t.x * m, t.y * m}
}

func PointInside[T any](p Point, a [][]T) bool {
	if p.x < 0 {
		return false
	}
	if p.x >= len(a) {
		return false
	}
	if p.y < 0 {
		return false
	}
	if p.y >= len(a[p.x]) {
		return false
	}
	return true
}

func PointInsideStrings(p Point, a []string) bool {
	if p.x < 0 {
		return false
	}
	if p.x >= len(a) {
		return false
	}
	if p.y < 0 {
		return false
	}
	if p.y >= len(a[p.x]) {
		return false
	}
	return true
}
