package kata

// DirReduc https://www.codewars.com/kata/550f22f4d758534c1100025a/train/go
func DirReduc(arr []string) []string {
	directions := Directions{}

	for _, direction := range arr {
		directions = directions.append(direction)
	}

	return directions.items
}

type Directions struct {
	items []string
}

func (ctx Directions) last() string {
	lastIndex := len(ctx.items) - 1
	return ctx.items[lastIndex]
}

func (ctx Directions) init() Directions {
	lastIndex := len(ctx.items) - 1
	return Directions{ctx.items[:lastIndex]}
}

func (ctx Directions) append(direction string) Directions {
	if ctx.isEmpty() {
		return Directions{append(ctx.items, direction)}
	}

	if areOpposites(ctx.last(), direction) {
		return ctx.init()
	}

	return Directions{append(ctx.items, direction)}
}

func (ctx Directions) isEmpty() bool {
	return len(ctx.items) == 0
}

func areOpposites(this string, other string) bool {
	opposites := map[string]string{
		"NORTH": "SOUTH",
		"SOUTH": "NORTH",
		"EAST":  "WEST",
		"WEST":  "EAST",
	}

	return opposites[this] == other
}

// Again what learned:
//   - define an own type like `type stack []string` and
//   - write context functions for it: `func (ctx stack) push(s string) stack { ...`
