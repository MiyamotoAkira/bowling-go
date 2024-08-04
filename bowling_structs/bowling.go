package bowling_structs

const maxFrameScore = 10
const firstRoll = 0
const secondRoll = 1
const thirdRoll = 2

type roll interface {
	score_pins(score int)
	is_scored() bool
}

type standardRoll struct {
	score  int
	scored bool
}

func (roll *standardRoll) is_scored() bool {
	return roll.scored
}

func (roll *standardRoll) score_pins(score int) {
	roll.score = score
	roll.scored = true
}

type frame interface {
	roll_pins(score int)
	is_full_frame() bool
	is_strike() bool
	is_spare() bool
	frame_score() int
}

type standardFrame struct {
	rolls [2]standardRoll
}

type finalFrame struct {
	rolls [3]standardRoll
}

func (frame *standardFrame) is_full_frame() bool {
	return frame.is_strike() || (frame.rolls[firstRoll].is_scored() && frame.rolls[secondRoll].is_scored())
}

func (frame *finalFrame) is_full_frame() bool {
	return frame.is_strike() || (frame.rolls[firstRoll].is_scored() && frame.rolls[secondRoll].is_scored())
}

func (frame *standardFrame) is_strike() bool {
	return frame.rolls[0].score == maxFrameScore
}

func (frame *finalFrame) is_strike() bool {
	return frame.rolls[0].score == maxFrameScore
}

func (frame *standardFrame) is_spare() bool {
	return !frame.is_strike() && (frame.rolls[0].score+frame.rolls[1].score == maxFrameScore)
}

func (frame *finalFrame) is_spare() bool {
	return !frame.is_strike() && (frame.rolls[0].score+frame.rolls[1].score == maxFrameScore)
}

func (frame *standardFrame) roll_pins(score int) {
	var roll_number = firstRoll
	if frame.rolls[firstRoll].is_scored() {
		roll_number = secondRoll
	}

	frame.rolls[roll_number].score_pins(score)
}

func (frame *finalFrame) roll_pins(score int) {
	var roll_number = firstRoll
	if frame.rolls[secondRoll].is_scored() {
		roll_number = thirdRoll
	} else if frame.rolls[firstRoll].is_scored() {
		roll_number = secondRoll
	}

	frame.rolls[roll_number].score_pins(score)
}

func (frame *standardFrame) frame_score() int {
	return frame.rolls[0].score + frame.rolls[1].score
}

func (frame *finalFrame) frame_score() int {
	return frame.rolls[firstRoll].score + frame.rolls[secondRoll].score + frame.rolls[thirdRoll].score
}

func (frame *finalFrame) two_roll_score() int {
	return frame.rolls[firstRoll].score + frame.rolls[secondRoll].score
}

type bowlingGame interface {
	roll_pins(score int)
	score_game() int
	start_game()
}

type game struct {
	frames        [9]standardFrame
	final_frame   finalFrame
	current_frame int
}

func (bowling_game *game) start_game() {
	bowling_game.frames = [9]standardFrame{}
	bowling_game.final_frame = finalFrame{}
	bowling_game.current_frame = 0
}

func (bowling_game *game) roll_pins(score int) {

	if bowling_game.current_frame > 8 {
		bowling_game.final_frame.roll_pins(score)
	} else {
		bowling_game.frames[bowling_game.current_frame].roll_pins(score)

		if bowling_game.frames[bowling_game.current_frame].is_full_frame() {
			bowling_game.current_frame++
		}
	}
}

func (bowling_game *game) score_game() int {
	total := 0
	for i, v := range bowling_game.frames {
		total += v.frame_score()
		if v.is_strike() {
			if i == 8 {
				total += bowling_game.final_frame.two_roll_score()
			} else if i == 7 {
				total += bowling_game.frames[i+1].frame_score()

				if bowling_game.frames[i+1].is_strike() {
					total += bowling_game.final_frame.rolls[0].score
				}
			} else {
				total += bowling_game.frames[i+1].frame_score()

				if bowling_game.frames[i+1].is_strike() {
					total += bowling_game.frames[i+2].rolls[0].score
				}
			}
		}

		if v.is_spare() {
			if i == 8 {
				total += bowling_game.final_frame.rolls[firstRoll].score
			} else {
				total += bowling_game.frames[i+1].rolls[firstRoll].score
			}
		}
	}

	total += bowling_game.final_frame.frame_score()

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
