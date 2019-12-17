package common

type Point struct {
	X int
	Y int
}

func (this Point) Min(other Point) Point {
	return Point{
		X: min(this.X, other.X),
		Y: min(this.Y, other.Y),
	}
}

func (this Point) Max(that Point) Point {
	return Point{
		X: max(this.X, that.X),
		Y: max(this.Y, that.Y),
	}
}

func min(x, y int) int {
	if y < x {
		return y
	}
	return x
}

func max(x, y int) int {
	if y > x {
		return y
	}
	return x
}