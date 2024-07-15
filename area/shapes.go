package area

// Square represents a square shape.
type Square struct {
	Side float64 // Length of the side of the square
}

// Area calculates the area of a square.
func (s Square) Area() float64 {
	return s.Side * s.Side
}
