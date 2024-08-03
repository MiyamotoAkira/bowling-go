package bowling

var scores [21]int
var current int

func Roll(roll int) {
	scores[current] = roll
	current++
}

func Score() int {
	result := 0
	for i, v := range scores {
		result += v
		if v == 10 {
			result += scores[i+1]
			result += scores[i+2]
		}
	}
	return result
}

func StartNewGame() {
	for i := range scores {
		scores[i] = 0
	}
	current = 0
}
