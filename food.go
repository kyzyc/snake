package main

type Food struct {
	Location Point
	Value    int
}

func CreateFood(x, y int) *Food {
	return &Food{Location: Point{x, y}, Value: 100}
}
