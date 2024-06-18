package kata

import (
	"regexp"
	"strconv"
)

// IsValidCoordinates / Coordinates Validator
// https://www.codewars.com/kata/5269452810342858ec000951/train/go
func IsValidCoordinates(coordinates string) bool {
	coords, err := NewCoordinates(coordinates)

	if err != nil {
		return false
	}

	return coords.IsValid()
}

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

func NewCoordinates(coordinates string) (Coordinates, error) {
	pattern, err := regexp.Compile(`^([-+]?([1-8]?\d(\.\d+)?|90(\.0+)?)),\s*([-+]?(180(\.0+)?|((1[0-7]\d)|([1-9]?\d))(\.\d+)?))$`)

	if err != nil {
		return Coordinates{}, err
	}

	if !pattern.MatchString(coordinates) {
		return Coordinates{}, nil
	}

	// parse groups
	groups := pattern.FindStringSubmatch(coordinates)

	println(groups[0])
	println(groups[1])
	println(groups[2])

	//latitudeAndLongitude := strings.Split(coordinates, ", ")

	latitude, err := strconv.ParseFloat(groups[1], 64)
	if err != nil {
		return Coordinates{}, err
	}

	longitude, err := strconv.ParseFloat(groups[2], 64)
	if err != nil {
		return Coordinates{}, err
	}

	return Coordinates{Latitude: latitude, Longitude: longitude}, nil
}

func (c Coordinates) IsValid() bool {
	return c.Latitude >= -90 &&
		c.Latitude <= 90 &&
		c.Longitude >= -180 &&
		c.Longitude <= 180
}
