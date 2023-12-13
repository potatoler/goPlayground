package pointerface

type HaveArea interface {
	Area() float32
}

type HavePerimeter interface {
	Perimeter() float32
}

type Square struct {
	side float32
}

func NewSquare(side float32) Square {
	return Square{side}
}

func (square *Square) Area() float32 {
	return square.side * square.side
}

func (square *Square) Perimeter() float32 {
	return 4 * square.side
}

type Rectangle struct {
	length, width float32
}

func (rect *Rectangle) Area() float32 {
	return rect.length * rect.width
}

type Circle struct {
	radius float32
}

func (circle *Circle) Area() float32 {
	return 3.14 * circle.radius * circle.radius
}

type Triangle struct {
	base, height float32
}

func (triangle *Triangle) Area() float32 {
	return 0.5 * triangle.base * triangle.height
}
