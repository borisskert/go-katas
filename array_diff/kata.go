package kata

// ArrayDiff https://www.codewars.com/kata/523f5d21c841566fde000009/train/go
func ArrayDiff(a, b []int) []int {
	bSet := newSetOf(b)

	var diff []int

	for _, v := range a {
		if _, ok := bSet[v]; !ok {
			diff = append(diff, v)
		}
	}

	return diff
}

func newSetOf(a []int) map[int]struct{} {
	set := make(map[int]struct{})

	for _, v := range a {
		set[v] = struct{}{}
	}

	return set
}
