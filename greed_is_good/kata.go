package kata

// Score / https://www.codewars.com/kata/5270d0d18625160ada0000e4/train/go
func Score(dice [5]int) int {
	scoreboard := newScoreboard()
	return scoreboard.scoreOf(dice)
}

type scores struct {
	scoreForThree int
	scoreForOne   int
}

type scoreboard struct {
	scoreboard map[int]scores
}

func newScoreboard() scoreboard {
	scoresMap := map[int]scores{}

	scoresMap[1] = scores{1000, 100}
	scoresMap[2] = scores{200, 0}
	scoresMap[3] = scores{300, 0}
	scoresMap[4] = scores{400, 0}
	scoresMap[5] = scores{500, 50}
	scoresMap[6] = scores{600, 0}

	return scoreboard{scoresMap}
}

func (ctx scoreboard) scoreOf(dice [5]int) int {
	score := 0
	groups := group(dice)

	for d, c := range groups {
		score += ctx.scoreFor(d, c)
	}

	return score
}

func (ctx scoreboard) scoreFor(dice, count int) int {
	p := ctx.scoreboard[dice]

	if count >= 3 {
		return p.scoreForThree + (count-3)*p.scoreForOne
	}

	return count * p.scoreForOne
}

func group(dice [5]int) map[int]int {
	groups := make(map[int]int)

	for _, d := range dice {
		groups[d]++
	}

	return groups
}
