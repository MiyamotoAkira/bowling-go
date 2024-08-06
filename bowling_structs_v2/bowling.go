package bowling_structs_v2

const maxFrameScore = 10
const firstRoll = 0
const secondRoll = 1
const thirdRoll = 2
const maxFrames = 10

type standardRoll struct {
	score  int
	scored bool
}

func (roll *standardRoll) score_pins(score int) {
	roll.score = score
	roll.scored = true
}

type standardFrame struct {
	rolls          [3]standardRoll
	is_final_frame bool
	last_roll      standardRoll
	bonus_score    [2]standardRoll
}

func (frame *standardFrame) is_full_frame() bool {
	if frame.is_final_frame {
		return frame.rolls[thirdRoll].scored
	} else {
		return frame.is_strike() || frame.rolls[secondRoll].scored
	}
}

func (frame *standardFrame) is_strike() bool {
	return frame.rolls[firstRoll].score == maxFrameScore
}

func (frame *standardFrame) is_spare() bool {
	return !frame.is_strike() && (frame.rolls[firstRoll].score+frame.rolls[secondRoll].score == maxFrameScore)
}

func (frame *standardFrame) roll_pins(score int) {
	var roll_number = firstRoll
	if frame.is_final_frame {
		if frame.rolls[secondRoll].scored {
			roll_number = thirdRoll
		} else if frame.rolls[firstRoll].scored {
			roll_number = secondRoll
		}

	} else {
		if frame.rolls[firstRoll].scored {
			roll_number = secondRoll
		}
	}

	frame.rolls[roll_number].score_pins(score)
}

func (frame *standardFrame) score() int {
	total := 0
	total += frame.rolls[firstRoll].score
	total += frame.rolls[secondRoll].score
	if frame.is_final_frame {
		total += frame.rolls[thirdRoll].score
	}
	total += frame.bonus_score[firstRoll].score
	total += frame.bonus_score[secondRoll].score

	return total
}

func (frame *standardFrame) add_bonus(score int) {
	if frame.is_strike() {
		if !frame.bonus_score[firstRoll].scored {
			frame.bonus_score[firstRoll].score_pins(score)
		} else if !frame.bonus_score[secondRoll].scored {
			frame.bonus_score[secondRoll].score_pins(score)
		}
	}

	if frame.is_spare() {
		if !frame.bonus_score[firstRoll].scored {
			frame.bonus_score[firstRoll].score_pins(score)
		}
	}
}

type game struct {
	frames        [10]standardFrame
	current_frame int
}

func (bowling_game *game) start_game() {
	bowling_game.frames = [maxFrames]standardFrame{}
	bowling_game.frames[maxFrames-1].is_final_frame = true
	bowling_game.current_frame = 0
}

func (bowling_game *game) roll_pins(score int) {
	bowling_game.frames[bowling_game.current_frame].roll_pins(score)

	if bowling_game.current_frame-1 >= 0 {
		bowling_game.frames[bowling_game.current_frame-1].add_bonus(score)
	}

	if bowling_game.current_frame-2 >= 0 {
		bowling_game.frames[bowling_game.current_frame-2].add_bonus(score)
	}

	if bowling_game.frames[bowling_game.current_frame].is_full_frame() {
		bowling_game.current_frame++
	}
}

func (bowling_game *game) score_game() int {
	total := 0
	for _, v := range bowling_game.frames {
		total += v.score()
	}

	return total
}

func Roll(score int) {
	bowling_game.roll_pins(score)
}

func Score() int {
	return bowling_game.score_game()
}

var bowling_game game

func StartNewGame() {
	bowling_game.start_game()
}
