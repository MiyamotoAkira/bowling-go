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
	frame_score_v2(nextFrame frame, nextNextFrame frame) int
	first_roll_score() int
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
	return frame.rolls[firstRoll].score == maxFrameScore
}

func (frame *finalFrame) is_strike() bool {
	return frame.rolls[firstRoll].score == maxFrameScore
}

func (frame *finalFrame) first_roll_score() int {
	return frame.rolls[firstRoll].score
}

func (frame *standardFrame) first_roll_score() int {
	return frame.rolls[firstRoll].score
}

func (frame *standardFrame) is_spare() bool {
	return !frame.is_strike() && (frame.rolls[firstRoll].score+frame.rolls[secondRoll].score == maxFrameScore)
}

func (frame *finalFrame) is_spare() bool {
	return !frame.is_strike() && (frame.rolls[firstRoll].score+frame.rolls[secondRoll].score == maxFrameScore)
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
	return frame.rolls[firstRoll].score + frame.rolls[secondRoll].score
}

func (frame *finalFrame) frame_score() int {
	return frame.rolls[firstRoll].score + frame.rolls[secondRoll].score
}

func (frame *finalFrame) full_frame_score() int {
	return frame.rolls[firstRoll].score + frame.rolls[secondRoll].score + frame.rolls[thirdRoll].score
}

func (frame *standardFrame) frame_score_v2(nextFrame frame, nextNextFrame frame) int {
	total := frame.rolls[firstRoll].score + frame.rolls[secondRoll].score
	if frame.is_strike() {
		if nextFrame != nil {
			total += nextFrame.frame_score()

			if nextFrame.is_strike() && nextNextFrame != nil {
				total += nextNextFrame.first_roll_score()
			}
		}
	}

	if frame.is_spare() && nextFrame != nil {
		total += nextFrame.first_roll_score()
	}

	return total
}

func (frame *finalFrame) frame_score_v2(nextFrame frame, nextNextFrame frame) int {
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

		var nextFrame frame = nil
		var nextNextFrame frame = nil
		if i == 8 {
			nextFrame = &bowling_game.final_frame
		} else if i == 7 {
			nextFrame = &bowling_game.frames[i+1]
			nextNextFrame = &bowling_game.final_frame

		} else {
			nextFrame = &bowling_game.frames[i+1]
			nextNextFrame = &bowling_game.frames[i+2]
		}
		total += v.frame_score_v2(nextFrame, nextNextFrame)
	}

	total += bowling_game.final_frame.full_frame_score()

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
