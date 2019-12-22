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
		X: v.X + o.X,
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

func (v Vector) Sign() Vector {
	return Vector{
		X: Sign(v.X),
		Y: Sign(v.Y),
		Z: Sign(v.Z),
	}
}

func (v Vector) Abs() int {
	return Abs(v.X) + Abs(v.Y) + Abs(v.Z)
}

func Sign(x int) int {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return 0
}

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Lcm(a, b int) int {
	return a / Gcd(a, b) * b
}