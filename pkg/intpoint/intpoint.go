package intpoint

import bp "github.com/dsabdrashitov/adventofcode2024/pkg/boilerplate"

type Point struct {
	X int
	Y int
}

var UP = Point{-1, 0}
var DOWN = Point{1, 0}
var LEFT = Point{0, -1}
var RIGHT = Point{0, 1}

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

var DIR9 = []Point{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 0},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

var DIRDIAG = []Point{
	{-1, 1},
	{1, 1},
	{1, -1},
	{-1, -1},
}

var DIRDR = []Point{
	{0, 0},
	{0, 1},
	{1, 0},
	{1, 1},
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

func (t Point) Compare(o Point) int {
	if t.X < o.X {
		return bp.LESS
	}
	if t.X > o.X {
		return bp.GREATER
	}
	if t.Y < o.Y {
		return bp.LESS
	}
	if t.Y > o.Y {
		return bp.GREATER
	}
	return bp.EQUAL
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
