package main

type Food struct {
	Location Point
	Value    int
}

//type FoodList struct {
//	Foods []Food
//}

func CreateFood(x, y int) *Food {
	return &Food{Location: Point{x, y}, Value: 100}
}
