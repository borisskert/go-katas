package kata

// SmallestIntegerFinder https://www.codewars.com/kata/55a2d7ebe362935a210000b2/train/go
func SmallestIntegerFinder(numbers []int) int {
	smallest := numbers[0]

	for _, number := range numbers {
		if number < smallest {
			smallest = number
		}
	}

	return smallest
}
