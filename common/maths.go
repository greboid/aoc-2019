package common

type Vector struct {
	X int
	Y int
	Z int
}

type Point struct {
	X int
	Y int
}

func (v Vector) Plus(o Vector) Vector {
	return Vector{
		X: v.Z + o.X,
		Y: v.Y + o.Y,
		Z: v.Z + o.Z,
	}
}

func (v Vector) Minus(o Vector) Vector {
	return Vector{
		X: v.X - o.X,
		Y: v.Y - o.Y,
		Z: v.Z - o.Z,
	}
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
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