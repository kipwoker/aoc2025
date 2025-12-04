package solutions

type Solution interface {
	Day() string
	Execute1(input string) string
	Execute2(input string) string
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sign(x int) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}

type Point struct {
	X, Y int
}

func (p Point) Hash() uint64 {
	return uint64(p.X)<<32 | uint64(p.Y)
}

func (p Point) Equals(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}

func (p Point) GetNeighbors8() []Point {
	return []Point{
		{p.X - 1, p.Y - 1},
		{p.X - 1, p.Y + 1},
		{p.X - 1, p.Y},
		{p.X, p.Y - 1},
		{p.X, p.Y + 1},
		{p.X + 1, p.Y - 1},
		{p.X + 1, p.Y},
		{p.X + 1, p.Y + 1},
	}
}
