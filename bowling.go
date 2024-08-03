package bowling

var scores [21]int
var current int

func Roll(roll int) {
	scores[current] = roll
	current++
}

func Score() int {
	result := 0
	newFrame := true
	numberOfFrames := 1
	for i, v := range scores {
		if numberOfFrames > 10 {
			break
		}
		if v == -1 {
			newFrame = !newFrame
			if newFrame {
				numberOfFrames++
			}
			continue
		}

		result += v
		if v == 10 {
			if scores[i+1] != -1 {
				result += scores[i+1]
			}
			if scores[i+2] != -1 {
				result += scores[i+2]
			}
			newFrame = true
		} else {
			if !newFrame {
				if v+scores[i-1] == 10 {
					if scores[i+1] != -1 {
						result += scores[i+1]
					}
				}
			}
			newFrame = !newFrame
		}
		if newFrame {
			numberOfFrames++
		}
	}

	return result
}

func StartNewGame() {
	for i := range scores {
		scores[i] = -1
	}
	current = 0
}
