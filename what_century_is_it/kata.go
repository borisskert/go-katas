package kata

import (
	"fmt"
	"strconv"
)

// WhatCentury / What century is it?
// https://www.codewars.com/kata/52fb87703c1351ebd200081f/train/go
func WhatCentury(yearString string) string {
	year := parseYear(yearString)

	century := (year-1)/100 + 1

	return toRank(century)
}

func parseYear(year string) int {
	yearInt, err := strconv.Atoi(year)

	if err != nil {
		panic(err)
	}

	return yearInt
}

func toRank(century int) string {
	if century%10 == 1 && century != 11 {
		return fmt.Sprint(century) + "st"
	}

	if century%10 == 2 && century != 12 {
		return fmt.Sprint(century) + "nd"
	}

	if century%10 == 3 && century != 13 {
		return fmt.Sprint(century) + "rd"
	}

	return fmt.Sprint(century) + "th"
}
