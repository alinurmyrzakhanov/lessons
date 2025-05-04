package physics

import (
	"fmt"

	"github.com/alinurmyrzakhanov/lessons/lesson7/example1/math"
	"github.com/alinurmyrzakhanov/lessons/lesson7/example1/physics/cool"
)

type Physic struct {
	Name string
}

func SecondLawOfNewton() int {
	m := &math.Math{
		Name: "Euler",
	}
	fmt.Println(m)
	return internal.Force(3, 4)
}
