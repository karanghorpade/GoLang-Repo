package main

import (
"fmt"
)

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	a := truck{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		fourWheel: true,
	}

	b := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "blue",
		},
		luxury: true,
	}

	fmt.Println(a.vehicle, a.fourWheel)

	fmt.Println(b.vehicle, b.luxury)
}

