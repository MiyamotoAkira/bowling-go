package bowling_simple

// 21 represents the maximum amount of possible rolls that could happen
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
		// We have already calculated all possible frames
		if numberOfFrames > 10 {
			break
		}

		// if we found the sentinel value no more rolls
		// have been made
		if v == -1 {
			break
		}

		// Generally we add the current value
		result += v
		if v == 10 {
			// This branch represents strikes
			if scores[i+1] != -1 {
				result += scores[i+1]
			}
			if scores[i+2] != -1 {
				result += scores[i+2]
			}
			newFrame = true
		} else {
			if !newFrame {
				// You can only see if is spare on the second roll
				// of the frame
				if v+scores[i-1] == 10 {
					// this branch represents spares
					if scores[i+1] != -1 {
						result += scores[i+1]
					}
				}
			}

			// We toggle the beginning of the frame marker
			newFrame = !newFrame
		}

		// We increment the current frame number
		if newFrame {
			numberOfFrames++
		}
	}

	return result
}

func StartNewGame() {
	for i := range scores {
		// -1 represents a sentinel value
		// No roll has been made yet
		scores[i] = -1
	}
	current = 0
}
