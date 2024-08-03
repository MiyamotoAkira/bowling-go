package bowling

var scores [21]int
var current int

func Roll(roll int) {
	scores[current] = roll
	current++
}

func Score() int {
	result := 0
	for _, v := range scores {
		result += v
	}
	return result
}

func Clear() {
	for i := range scores {
		scores[i] = 0
	}
	current = 0
}
