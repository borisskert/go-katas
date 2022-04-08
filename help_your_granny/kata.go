package kata

import "math"

// Tour https://www.codewars.com/kata/5536a85b6ed4ee5a78000035/train/go
func Tour(arrFriends []string, ftwns map[string]string, h map[string]float64) int {
	friends := Strings{arrFriends}
	friendTowns := FriendTowns{ftwns}
	towns := friendTowns.Filter(friends)
	townDistances := TownDistances{h}
	distances := townDistances.Filter(towns)

	distance := GranniesDistance(distances)

	return int(math.Floor(distance))
}

func GranniesDistance(distances Distances) float64 {
	var distance = 0.0

	var a float64
	var c float64

	for index, item := range distances.items {
		if index == 0 {
			c = item
			continue
		}

		a = c
		c = item

		distance += OppositeLeg(a, c)
	}

	return distances.Head() + distance + distances.Last()
}

// OppositeLeg calculates the opposite leg [MATH.] (c) by the Pythagoras' theorem
// of the specified adjacent leg [MATH.] (a) and the hypotenuse [MATH.] (c)
func OppositeLeg(a float64, c float64) float64 {
	return math.Sqrt(c*c - a*a)
}

type Strings struct {
	items []string
}

func (x Strings) Contains(element string) bool {
	for _, s := range x.items {
		if s == element {
			return true
		}
	}

	return false
}

type FriendTowns struct {
	items map[string]string
}

func (x FriendTowns) Contains(friend string) bool {
	for k := range x.items {
		if k == friend {
			return true
		}
	}

	return false
}

func (x FriendTowns) Filter(friends Strings) Strings {
	var result []string

	for _, friend := range friends.items {
		if x.Contains(friend) {
			result = append(result, x.items[friend])
		}
	}

	return Strings{result}
}

type TownDistances struct {
	items map[string]float64
}

func (x TownDistances) Contains(town string) bool {
	for k := range x.items {
		if k == town {
			return true
		}
	}

	return false
}

type Distances struct {
	items []float64
}

func (x TownDistances) Filter(towns Strings) Distances {
	var result []float64

	for _, town := range towns.items {
		if x.Contains(town) {
			result = append(result, x.items[town])
		}
	}

	return Distances{result}
}

func (x Distances) Head() float64 {
	return x.items[0]
}

func (x Distances) Last() float64 {
	return x.items[len(x.items)-1]
}
